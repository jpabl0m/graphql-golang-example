package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

type contextKey string

const userIDContextKey contextKey = "userID"

func authenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secretKey := os.Getenv("JWT_SECRET_KEY")
		token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(secretKey), nil
		})
		if err != nil || !token.Valid {
			log.Print(err)
			c.JSON(http.StatusForbidden, gin.H{"error": "token invalid"})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && claims["sub"] != nil {
			ctx := context.WithValue(c.Request.Context(), userIDContextKey, claims["sub"])
			c.Request = c.Request.WithContext(ctx)
			c.Next()
			return
		}
		c.JSON(http.StatusForbidden, gin.H{"error": "claims insufficient"})
		c.Abort()
		return
	}

}
