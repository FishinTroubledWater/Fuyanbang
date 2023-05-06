package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type MyClaims struct {
	Account            string
	PhoneNumber        string
	jwt.StandardClaims // 标准Claims结构体，可设置8个标准字段
}

var (
	secret         = []byte("31231dasdaseqwkjcozx") //秘钥
	ExpireDuration = 3 * time.Minute                //秘钥有效时间
	expirationTime = time.Now().Add(ExpireDuration)
	issuer         = "Fyb" //签发人
)

func GenerateToken(account string, phoneNumber string) (string, error) {
	claims := &MyClaims{
		Account:     account,
		PhoneNumber: phoneNumber,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tokenString, err := token.SignedString(secret); err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func JwtVerify(c *gin.Context) error {
	token := c.GetHeader("token")
	if token == "" {
		return errors.New("token not exists")
	}
	myClaims, err1 := parseToken(token)
	c.Set("claims", myClaims)
	return err1
}

func parseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token can't be parsed or is not valid")
}
