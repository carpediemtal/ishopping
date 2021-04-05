package main

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
)

var identityKey = "identity"

func initAuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "gin-jwt",
		Key:         []byte("secret key"),
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{username: claims[identityKey].(string)}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			type login struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}
			var loginVals login
			if err := c.ShouldBindJSON(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			user, err := getUserByUsername(loginVals.Username)
			if err != nil || user.password != loginVals.Password {
				return nil, jwt.ErrFailedAuthentication
			}
			return &user, nil
		},
		Unauthorized: func(c *gin.Context, code int, msg string) {
			JsonErr(c, msg)
		},
	})
	if err != nil {
		log.Fatalf("JWT Error: %s\n", err.Error())
	}
	return authMiddleware
}
