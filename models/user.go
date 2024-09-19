package models

import (
	"fmt"
	"friends_ranking/config/dbconn"
	"friends_ranking/config/globalConst"
	"friends_ranking/config/variable"
	"friends_ranking/service/user/user_cach_redis"
	"time"

	"go.uber.org/zap"
)

type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	UserName string `json:"userName" gorm:"column:userName"`
	NickName string `json:"nickName" gorm:"column:nickName"`
	Account  string `json:"account"  gorm:"column:account"`
	Password string `json:"passWord" gorm:"column:passWord"`
	Email    string `json:"email" gorm:"column:email" `
	Age      int    `json:"age"  gorm:"column:age" `
	Phone    string `json:"phone" gorm:"column:phone"`
	Job      string `json:"job" gorm:"column:job"`
	Level    int    `json:"level"  gorm:"column:level"`
	BirthDay string `json:"birthDay"  gorm:"column:birthDay"`
	Sex      int    `json:"sex"  gorm:"column:sex"`
	IsDelete int    `json:"isDelete" gorm:"column:isDelete"`
	IdCard   string `json:"idCard" gorm:"column:idCard"`

	Status        int    `json:"status" gorm:"column:status"`
	RegisterTime  string `json:"registerTime" gorm:"column:registerTime"`
	LastLoginTime string `json:"lastLoginTime" gorm:"column:lastLoginTime"`
	CreateTime    string `json:"createTime" gorm:"column:createTime"`
	EditTime      string `json:"editTime" gorm:"column:editTime"`
	Creator       string `json:"creator" gorm:"column:creator"`
	Editor        string `json:"editor" gorm:"column:editor"`
	dbconn.BaseModel
}

func (u *User) TableName() string {
	return "sys_user"
}

func CreateUserFactory(sqlType string) *User {
	return &User{BaseModel: dbconn.BaseModel{
		DB:        dbconn.UseDbConn(sqlType),
		Id:        0,
		CreatedAt: "",
		UpdatedAt: "",
	}}
}

/*
*
根据用户账号查询用户是否存在
*/
func (u *User) SelectByAccount(account string) (*User, error) {
	sql := "select * from sys_user where account=? and isDelete=0 "
	rersult := u.Raw(sql, account).First(u)
	if rersult.Error == nil {
		return u, nil
	} else {
		return nil, rersult.Error
	}
}

func (u *User) InsertUser(u2 *User) bool {
	sql := `
		insert into sys_user (userName, nickName, account, password, idCard, email, age, phone, job, level, birthDay, status, registerTime,  creator,  editTor, isDelete) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`
	fmt.Println("u2 info: ", u2)
	if u.Exec(sql, u2.UserName, u2.NickName, u2.Account, u2.Password, u2.IdCard, u2.Email, u2.Age, u2.Phone, u2.Job, u2.Level, u2.BirthDay, u2.Status, u2.RegisterTime, u2.Creator, u2.Editor, u2.IsDelete).Error == nil {
		return true
	}
	return false
}

func (u *User) UpdateUser(u2 *User) bool {
	err := u.DB.Model(&User{}).Where("id = ?", u2.Id).Updates(u2).Error
	if err != nil {
		return false
	}
	return true
}

// 记录用户登陆（login）生成的token，每次登陆记录一次token
func (u *User) OauthLoginToken(userId string, token string, expiresAt int64, clientIp string) bool {
	sql := `
		INSERT   INTO  tb_oauth_access_tokens(fr_user_id,action_name,token,expires_at,client_ip)
		SELECT  ?,'login',? ,?,? FROM DUAL    WHERE   NOT   EXISTS(SELECT  1  FROM  tb_oauth_access_tokens a WHERE  a.fr_user_id=?  AND a.action_name='login' AND a.token=?  )
	`
	//注意：token的精确度为秒，如果在一秒之内，一个账号多次调用接口生成的token其实是相同的，这样写入数据库，第二次的影响行数为0，知己实际上操作仍然是有效的。
	//所以这里只判断无错误即可，判断影响行数的话，>=0 都是ok的
	if u.Exec(sql, userId, token, time.Unix(expiresAt, 0).Format(globalConst.DateFormat), clientIp, userId, token).Error == nil {
		// 异步缓存用户有效的token到redis
		if variable.YamlConfig.GetInt("Token.IsCacheToRedis") == 1 {
			go u.ValidTokenCacheToRedis(userId)
		}
		return true
	}
	return false
}

// 后续两个函数专门处理用户 token 缓存到 redis 逻辑

func (u *User) ValidTokenCacheToRedis(userId string) {
	tokenCacheRedisFact := user_cach_redis.CreateUsersTokenCacheFactory(userId)
	if tokenCacheRedisFact == nil {
		variable.ZapLog.Error("redis连接失败，请检查配置")
		return
	}
	defer tokenCacheRedisFact.ReleaseRedisConn()

	sql := "SELECT   token,expires_at  FROM  `tb_oauth_access_tokens`  WHERE   fr_user_id=?  AND  revoked=0  AND  expires_at>NOW() ORDER  BY  expires_at  DESC , updated_at  DESC  LIMIT ?"
	maxOnlineUsers := variable.YamlConfig.GetInt("Token.JwtTokenOnlineUsers")
	rows, err := u.Raw(sql, userId, maxOnlineUsers).Rows()
	defer func() {
		//  凡是获取原生结果集的查询，记得释放记录集
		_ = rows.Close()
	}()

	var tempToken, expires string
	if err == nil && rows != nil {
		for i := 1; rows.Next(); i++ {
			err = rows.Scan(&tempToken, &expires)
			if err == nil {
				if ts, err := time.ParseInLocation(globalConst.DateFormat, expires, time.Local); err == nil {
					tokenCacheRedisFact.SetTokenCache(ts.Unix(), tempToken)
					// 因为每个用户的token是按照过期时间倒叙排列的，第一个是有效期最长的，将该用户的总键设置一个最大过期时间，到期则自动清理，避免不必要的数据残留
					if i == 1 {
						tokenCacheRedisFact.SetUserTokenExpire(ts.Unix())
					}
				} else {
					variable.ZapLog.Error("expires_at 转换位时间戳出错", zap.Error(err))
				}
			}
		}
	}
	// 缓存结束之后删除超过系统设置最大在线数量的token
	tokenCacheRedisFact.DelOverMaxOnlineCache()
}

// 用户刷新token,条件检查: 相关token在过期的时间之内，就符合刷新条件
func (u *User) OauthRefreshConditionCheck(userId string, oldToken string) bool {
	// 首先判断旧token在本系统自带的数据库已经存在，才允许继续执行刷新逻辑
	var oldTokenIsExists int
	sql := "SELECT count(*)  as  counts FROM tb_oauth_access_tokens  WHERE fr_user_id =? and token=? and NOW()<DATE_ADD(expires_at,INTERVAL  ? SECOND)"
	if u.Raw(sql, userId, oldToken, variable.YamlConfig.GetInt64("Token.JwtTokenRefreshAllowSec")).First(&oldTokenIsExists).Error == nil && oldTokenIsExists == 1 {
		return true
	}
	return false
}

// 用户刷新token
func (u *User) OauthRefreshToken(userId string, expiresAt int64, oldToken, newToken, clientIp string) bool {
	sql := "UPDATE   tb_oauth_access_tokens   SET  token=? ,expires_at=?,client_ip=?,updated_at=NOW(),action_name='refresh'  WHERE   fr_user_id=? AND token=?"
	if u.Exec(sql, newToken, time.Unix(expiresAt, 0).Format(globalConst.DateFormat), clientIp, userId, oldToken).Error == nil {
		// 异步缓存用户有效的token到redis
		if variable.YamlConfig.GetInt("Token.IsCacheToRedis") == 1 {
			go u.ValidTokenCacheToRedis(userId)
		}
		go u.UpdateUserloginInfo(clientIp, userId)
		return true
	}
	return false
}

// 更新用户登陆次数、最近一次登录ip、最近一次登录时间
func (u *User) UpdateUserloginInfo(last_login_ip string, userId string) {
	sql := "UPDATE  tb_users   SET  login_times=IFNULL(login_times,0)+1,last_login_ip=?,last_login_time=?  WHERE   id=?  "
	_ = u.Exec(sql, last_login_ip, time.Now().Format(globalConst.DateFormat), userId)
}

// 判断用户token是否在数据库存在+状态OK
func (u *User) OauthCheckTokenIsOk(userId string, token string) bool {
	sql := "SELECT   token  FROM  `tb_oauth_access_tokens`  WHERE   fr_user_id=?  AND  revoked=0  AND  expires_at>NOW() ORDER  BY  expires_at  DESC , updated_at  DESC  LIMIT ?"
	maxOnlineUsers := variable.YamlConfig.GetInt("Token.JwtTokenOnlineUsers")
	rows, err := u.Raw(sql, userId, maxOnlineUsers).Rows()
	defer func() {
		//  凡是查询类记得释放记录集
		_ = rows.Close()
	}()
	if err == nil && rows != nil {
		for rows.Next() {
			var tempToken string
			err := rows.Scan(&tempToken)
			if err == nil {
				if tempToken == token {
					return true
				}
			}
		}
	}
	return false
}
