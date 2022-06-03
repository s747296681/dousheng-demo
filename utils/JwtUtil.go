package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"time"
)

type CustomClaims struct {
	UserName string
	PassWord string
	jwt.StandardClaims
}

var MySecret = []byte("测试密钥")

// 创建 Token
func GenToken(userName string, passWord string) (string, error) {
	claim := CustomClaims{
		userName,
		passWord,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * 1)), //1小时后过期
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

// 解析 token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		fmt.Println(" token parse err:", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 刷新 Token
func RefreshToken(tokenStr string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = jwt.At(time.Now().Add(time.Minute * 10))
		return GenToken(claims.UserName, claims.PassWord)
	}
	return "", errors.New("Cloudn't handle this token")
}

func GetClaimInfoByCtx(c *gin.Context) (*CustomClaims, error) {
	claims, ok := c.Get("claims")
	if !ok {
		return nil, fmt.Errorf("could not get claims")
	}
	return claims.(*CustomClaims), nil
}
