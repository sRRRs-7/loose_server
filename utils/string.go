package utils

import (
	"strings"

	"github.com/go-redis/redis/v9"
)

func GetUsername(redisValue *redis.StringCmd) string {
	s := strings.Split(redisValue.String(), ",")
	s = strings.Split(s[1], ":")
	username := s[1]
	username = username[1:]
	username = username[:len(username)-1]

	return username
}
