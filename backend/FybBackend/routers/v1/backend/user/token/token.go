package token

import (
	fybDatabase "FybBackend/database"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type MyClaims struct {
	Admin              fybDatabase.Admin
	jwt.StandardClaims // 标准Claims结构体，可设置8个标准字段
}

var (
	//自定义的token秘钥
	secret = []byte("16849841325189456f487")
	//token有效时间（纳秒）
	ExpireDuration = 3 * time.Minute
)

func GenerateToken(adminInfo fybDatabase.Admin) (string, error) {
	expirationTime := time.Now().Add(ExpireDuration) // 两个小时有效期
	claims := &MyClaims{
		Admin: adminInfo,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "Fyb",
		},
	}
	// 生成Token，指定签名算法和claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名
	if tokenString, err := token.SignedString(secret); err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

// 验证token
func JwtVerify(c *gin.Context) error {
	token := c.GetHeader("token")
	if token == "" {
		return errors.New("token not exist")
	}
	//验证token，并存储在请求中
	myClaims, _ := parseToken(token)
	c.Set("admin", myClaims)
	return nil
}

// 解析Token
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
	return nil, errors.New("token无法解析")
}

// 更新token
func Refresh(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		panic("token is valid")
	}
	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(24 * time.Hour).Unix()
	return GenerateToken(claims.Admin)
}
