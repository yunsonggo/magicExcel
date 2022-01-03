package jwt

import (
	"2021/magicExcel/server/conf"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

// 发放token
func ReleaseToken(userId int64) (string, error) {
	jwtIssuer := conf.AppConf.Issuer
	jwtKey := []byte(conf.AppConf.JwtKey)
	expireTime := time.Duration(conf.AppConf.Expire*60) * time.Minute
	// 有效期
	expirationTime := time.Now().Add(expireTime)
	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			// 有效期
			ExpiresAt: expirationTime.Unix(),
			// 开始时间
			IssuedAt: time.Now().Unix(),
			// 发放者
			Issuer: jwtIssuer,
			// 主题
			Subject: "user token",
		},
	}
	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	jwtKey := []byte(conf.AppConf.JwtKey)
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
