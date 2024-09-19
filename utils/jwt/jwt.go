package jwy

import (
	"errors"
	"friends_ranking/config/errorMsg"
	"friends_ranking/config/globalConst"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 定义一个 JWT验签 结构体
type JwtSign struct {
	SigningKey []byte
}

// 自定义JWT payload结构体
type CustomClaims struct {
	Account string `json:"account"` //自定义用户类型，游客、内部用户、合作用户或者其他标识
	Status  int    `json:"status"`
	jwt.StandardClaims
}

// 使用工厂创建一个 JWT 结构体
func CreateMyJWT(signKey string) *JwtSign {
	if len(signKey) <= 0 {
		signKey = globalConst.JWTSecretKey
	}
	return &JwtSign{
		[]byte(signKey),
	}
}

// CreateToken 生成一个token
func (j *JwtSign) CreateToken(claims CustomClaims) (string, error) {
	// 生成jwt格式的header、claims 部分
	tokenPartA := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 继续添加秘钥值，生成最后一部分
	return tokenPartA.SignedString(j.SigningKey)
}

// 解析Token
func (j *JwtSign) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if token == nil {
		return nil, errors.New(errorMsg.ErrorsTokenInvalid)
	}
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(errorMsg.ErrorsTokenMalFormed)
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(errorMsg.ErrorsTokenNotActiveYet)
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// 如果 TokenExpired ,只是过期（格式都正确），我们认为他是有效的，接下可以允许刷新操作
				token.Valid = true
				goto labelHere
			} else {
				return nil, errors.New(errorMsg.ErrorsTokenInvalid)
			}
		}
	}
labelHere:
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New(errorMsg.ErrorsTokenInvalid)
	}
}

// 更新token
func (j *JwtSign) RefreshToken(tokenString string, extraAddSeconds int64) (string, error) {

	if CustomClaims, err := j.ParseToken(tokenString); err == nil {
		CustomClaims.ExpiresAt = time.Now().Unix() + extraAddSeconds
		return j.CreateToken(*CustomClaims)
	} else {
		return "", err
	}
}
