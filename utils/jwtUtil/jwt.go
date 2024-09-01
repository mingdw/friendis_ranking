package jwtutil

import (
	"errors"
	"fmt"
	"friends_ranking/config/errorMsg"
	"friends_ranking/config/globalConst"

	"github.com/dgrijalva/jwt-go"
)

// 定义一个 JWT验签 结构体
type JwtSign struct {
	SigningKey []byte
}

// 自定义JWT payload结构体
type CustomClaims struct {
	Type  string `json:"type"` //自定义用户类型，游客、内部用户、合作用户或者其他标识
	Phone string `json:"phone"`
	jwt.StandardClaims
}

func JWTInstance(key string) *JwtSign {
	fmt.Println("创建JWT TOKEN")
	if len(key) <= 0 {
		key = globalConst.JWTSecretKey
	}

	return &JwtSign{
		[]byte(key),
	}
}

// CreateToken 生成一个token
func (j *JwtSign) CreateToken(claims CustomClaims) (string, error) {
	// 生成jwt格式的header、claims 部分
	tokenPartA := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 继续添加秘钥值，生成最后一部分
	return tokenPartA.SignedString(j.SigningKey)
}

func (j *JwtSign) ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
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
