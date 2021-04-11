package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	UserID     int
	Username   string
	JWTVersion int
	jwt.StandardClaims
}

func NewJWT(c *CustomClaims) (token string, err error) {
	c.JWTVersion = 1
	// set expired time of jwt
	c.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		Issuer:    "ishopping",
	}
	// set signing method of jwt
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// generate jwt string
	token, err = tokenClaims.SignedString([]byte("B2VDXZ*oJsVLw13$$k29UPz*fm1Hp0B8jz6i7*1G7Y^9C#wQNHIO51J9DdyxJGobZj%Qobc@wNDIiEiK0kGN6QGd&ZFQ0wmV5k1"))
	return
}

func ParseJWT(token string) (*CustomClaims, error) {
	// parse the token string to custom claim struct
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("B2VDXZ*oJsVLw13$$k29UPz*fm1Hp0B8jz6i7*1G7Y^9C#wQNHIO51J9DdyxJGobZj%Qobc@wNDIiEiK0kGN6QGd&ZFQ0wmV5k1"), nil
	})
	if err != nil || tokenClaims == nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("failed to valid the jwt")
	}
}
