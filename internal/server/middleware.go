package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/stobita/graphql-golang-example/internal/usecase"
)

type middleware struct {
	sessionStore
}

type sessionStore interface {
	Get(sessionID string) (string, error)
}

func newMiddleware(s sessionStore) *middleware {
	return &middleware{
		s,
	}
}

const userSessionIDCookieName = "user_session"

func (m *middleware) getAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Request.Cookie(userSessionIDCookieName)
		if err != nil || c == nil {
			c.Next()
			return
		}
		userID, err := m.sessionStore.Get(sessionID.String())
		if err != nil {
			c.Next()
			return
		}
		ctx := context.WithValue(c.Request.Context(), usecase.UserIDContextKey, userID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
