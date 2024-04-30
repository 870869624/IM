package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	mySigningKey = "test"
)

type UserAuthClaims struct {
	Account string `json:"account"`
	jwt.RegisteredClaims
}

// 获取token
func CreatToken(account string) (string, error) {
	claims := UserAuthClaims{
		account,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(mySigningKey))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenString, nil
}

// 思路：传入tokenstring 获取claims里的account
func ParaseToken(tokenString string) (*UserAuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserAuthClaims{}, func(t *jwt.Token) (interface{}, error) {//注意这里的结构体如果不用指针就会在传入tokenstring的时候报错
		return []byte(mySigningKey), nil
	})
	if err != nil {
		return nil, err //后期替换为服务器内部错误
	}
	if claims, ok := token.Claims.(*UserAuthClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
