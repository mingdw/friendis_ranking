package user_token

import (
	"errors"
	"friends_ranking/config/errorMsg"
	"friends_ranking/config/globalConst"
	"friends_ranking/config/variable"
	"friends_ranking/models"
	token_cache_redis "friends_ranking/service/user/user_cach_redis"
	Myjwt "friends_ranking/utils/jwt"

	"github.com/dgrijalva/jwt-go"

	"time"
)

// CreateUserFactory 创建 userToken 工厂
func CreateUserFactory() *userToken {
	return &userToken{
		userJwt: Myjwt.CreateMyJWT(variable.YamlConfig.GetString("Token.JwtTokenSignKey")),
	}
}

type userToken struct {
	userJwt *Myjwt.JwtSign
}

// GenerateToken 生成token
func (u *userToken) GenerateToken(userid string, account string, status int, expireAt int64) (tokens string, err error) {

	// 根据实际业务自定义token需要包含的参数，生成token，注意：用户密码请勿包含在token
	customClaims := Myjwt.CustomClaims{
		Account: account,
		Status:  status,
		// 特别注意，针对前文的匿名结构体，初始化的时候必须指定键名，并且不带 jwt. 否则报错：Mixture of field: value and value initializers
		StandardClaims: jwt.StandardClaims{
			Id:        userid,
			NotBefore: time.Now().Unix() - 10,       // 生效开始时间
			ExpiresAt: time.Now().Unix() + expireAt, // 失效截止时间
		},
	}
	return u.userJwt.CreateToken(customClaims)
}

// RecordLoginToken 用户login成功，记录用户token
func (u *userToken) RecordLoginToken(userToken, clientIp string) bool {
	if customClaims, err := u.userJwt.ParseToken(userToken); err == nil {
		userId := customClaims.Id
		expiresAt := customClaims.ExpiresAt
		return models.CreateUserFactory("").OauthLoginToken(userId, userToken, expiresAt, clientIp)
	} else {
		return false
	}
}

// TokenIsMeetRefreshCondition 检查token是否满足刷新条件
func (u *userToken) TokenIsMeetRefreshCondition(token string) bool {
	// token基本信息是否有效：1.过期时间在允许的过期范围内;2.基本格式正确
	customClaims, code := u.isNotExpired(token, variable.YamlConfig.GetInt64("Token.JwtTokenRefreshAllowSec"))
	switch code {
	case globalConst.JwtTokenOK, globalConst.JwtTokenExpired:
		//在数据库的存储信息是否也符合过期刷新刷新条件
		if models.CreateUserFactory("").OauthRefreshConditionCheck(customClaims.Id, token) {
			return true
		}
	}
	return false
}

// RefreshToken 刷新token的有效期（默认+3600秒，参见常量配置项）
func (u *userToken) RefreshToken(oldToken, clientIp string) (newToken string, res bool) {
	var err error
	//如果token是有效的、或者在过期时间内，那么执行更新，换取新token
	if newToken, err = u.userJwt.RefreshToken(oldToken, variable.YamlConfig.GetInt64("Token.JwtTokenRefreshExpireAt")); err == nil {
		if customClaims, err := u.userJwt.ParseToken(newToken); err == nil {
			userId := customClaims.Id
			expiresAt := customClaims.ExpiresAt
			if models.CreateUserFactory("").OauthRefreshToken(userId, expiresAt, oldToken, newToken, clientIp) {
				return newToken, true
			}
		}
	}

	return "", false
}

// 判断token本身是否未过期，不校验token的本身的过期时间，token的过期时间为redis的zset当中scro代表的过期时间
// 参数解释：
// token： 待处理的token值
// expireAtSec： 过期时间延长的秒数，主要用于用户刷新token时，判断是否在延长的时间范围内，非刷新逻辑默认为0
func (u *userToken) isNotExpired(token string, expireAtSec int64) (*Myjwt.CustomClaims, int) {
	if customClaims, err := u.userJwt.ParseToken(token); err == nil {

		if time.Now().Unix()-(customClaims.ExpiresAt+expireAtSec) < 0 {
			// token有效
			return customClaims, globalConst.JwtTokenOK
		} else {
			// 过期的token
			return customClaims, globalConst.JwtTokenExpired
		}
	} else {
		// 无效的token
		return nil, globalConst.JwtTokenInvalid
	}
}

// IsEffective 判断token是否有效（未过期+数据库用户信息正常）
func (u *userToken) IsEffective(token string) bool {
	customClaims, code := u.isNotExpired(token, 0)
	if globalConst.JwtTokenOK == code {
		//1.首先在redis检测是否存在某个用户对应的有效token，如果存在就直接返回，不再继续查询mysql，否则最后查询mysql逻辑，确保万无一失
		if variable.YamlConfig.GetInt("Token.IsCacheToRedis") == 1 {
			tokenRedisFact := token_cache_redis.CreateUsersTokenCacheFactory(customClaims.Id)
			if tokenRedisFact != nil {
				defer tokenRedisFact.ReleaseRedisConn()
				if tokenRedisFact.TokenCacheIsExists(token) {
					return true
				}
			}
		}
	}
	return false
}

// ParseToken 将 token 解析为绑定时传递的参数
func (u *userToken) ParseToken(tokenStr string) (CustomClaims Myjwt.CustomClaims, err error) {
	if customClaims, err := u.userJwt.ParseToken(tokenStr); err == nil {
		return *customClaims, nil
	} else {
		return Myjwt.CustomClaims{}, errors.New(errorMsg.ErrorsParseTokenFail)
	}
}

// DestroyToken 销毁token，基本用不到，因为一个网站的用户退出都是直接关闭浏览器窗口，极少有户会点击“注销、退出”等按钮，销毁token其实无多大意义
func (u *userToken) DestroyToken() {

}
