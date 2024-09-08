package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"qingyu-wf/global"
	"time"
)

// 定义密钥
var jwtKey = []byte("qinyu")

// Claims 结构体用来定义 JWT 的载荷
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 设置令牌过期时间
	claims := &Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "qinyu",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseJWT 解析并验证 JWT
func ParseJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func JWTAuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(200, global.RespMsg(7, "Authorization异常"))
		c.Abort()
		return
	}
	claims, err := ParseJWT(token)
	if err != nil {
		return
	}
	result, err := global.Redis.Get(c, claims.UserID).Result()
	if err != nil {
		return
	}
	if result != token {
		c.JSON(200, global.RespMsg(1, "该账号已在其他地方登录"))
		return
	}
	c.Set("id", claims.UserID)
	c.Next()
}
