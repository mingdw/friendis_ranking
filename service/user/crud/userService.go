package userService

import (
	"lottery_annual/config/globalConst"
	"lottery_annual/config/variable"
	"lottery_annual/models"
	"lottery_annual/service/user/user_cach_redis"
	"lottery_annual/utils/md5_encrypt"
	"lottery_annual/utils/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Register(user *models.User, c *gin.Context) {
	//先检查注册的账号是否存在

	userModels := models.CreateUserFactory("")
	userName := user.UserName
	ruser, err := userModels.SelectByUserName(userName)
	if err != nil {
		response.ReturnFail(c, "根据账号查询出错")
		return
	} else {
		if ruser.Id == 0 {
			user.Password = md5_encrypt.Base64Md5(user.Password)
			isInsert := userModels.InsertUser(user)
			if !isInsert {
				response.ReturnFail(c, "新增用户出错")
				return
			}
		} else {
			response.ReturnFail(c, "账户已存在，请输入新账号")
			return
		}
	}
	response.ReturnDefaultOk(c)
}

// 用户登陆校验缓存token主要逻辑
func ValidAndCatchToken(user *models.User, c *gin.Context, token string) bool {
	//1.利用redis的zset数据结构设置用户的最大过期时间，校验是否过期判断对应的token的score分数即可，每次刷新用户key的过期时间以及token的分数，如果校验只需要校验当前token的user_id是否存在
	tokenCacheRedisFact := user_cach_redis.CreateUsersTokenCacheFactory(strconv.Itoa(user.Id))
	if tokenCacheRedisFact == nil {
		variable.ZapLog.Error("redis连接失败，请检查配置")
		return false
	}
	defer tokenCacheRedisFact.ReleaseRedisConn()
	t := time.Now()
	tokenExpiren := variable.YamlConfig.GetInt64("Token.JwtTokenCreatedExpireAt")
	if ts, err := time.ParseInLocation(globalConst.DateFormat, t.Add(time.Duration(tokenExpiren)*time.Second).Format("2006-01-02 15:04:05"), time.Local); err == nil {
		catchResult := tokenCacheRedisFact.SetTokenCache(ts.Unix(), token)
		// 因为每个用户的token是按照过期时间倒叙排列的，第一个是有效期最长的，将该用户的总键设置一个最大过期时间，到期则自动清理，避免不必要的数据残留
		if catchResult {
			tokenCacheRedisFact.SetUserTokenExpire(ts.Unix())
			//移除过期的其它当前登陆用户的token
			tokenCacheRedisFact.DelOverMaxOnlineCache()
			return true
		}
	} else {
		variable.ZapLog.Error("expires_at 转换位时间戳出错", zap.Error(err))
	}

	return false
}
