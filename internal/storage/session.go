package storage

import "github.com/go-redis/redis"

type SessionStore interface {
	Get(key string) (string, error)
}

type sessionStore struct {
	client *redis.Client
}

func NewSessionStore(c *redis.Client) SessionStore {
	return &sessionStore{c}
}

func (s *sessionStore) Get(key string) (string, error) {
	v, err := s.client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return v, nil
}
