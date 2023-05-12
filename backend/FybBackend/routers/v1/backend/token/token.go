package token

import (
	"FybBackend/database"
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
	ExpireDuration = 24 * 7 * time.Hour             //秘钥有效时间
	issuer         = "Fyb"                          //签发人
)

func GenerateToken(admin database.Admin) (string, error) {
	expirationTime := time.Now().Add(ExpireDuration)
	claims := &MyClaims{
		Account:     admin.Account,
		PhoneNumber: admin.PhoneNumber,
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
	token := c.GetHeader("Authorization")
	if token == "" {
		return errors.New("token not exists")
	}
	myClaims, err1 := parseToken(token)
	c.Set("claims", myClaims)
	return err1
}
func InitToken(c *gin.Context) error {
	token := c.GetHeader("Authorization")
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

//func refresh(tokenString string) string {
//	jwt.TimeFunc = func() time.Time {
//		return time.Unix(0, 0)
//	}
//	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return secret, nil
//	})
//	if err != nil {
//		panic(err)
//	}
//	claims, ok := token.Claims.(*MyClaims)
//	if !ok {
//		panic("token is valid")
//	}
//	jwt.TimeFunc = time.Now
//	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
//	tokenStr, _ := GenerateToken(MyClaims.Pho)
//	return token
//}
