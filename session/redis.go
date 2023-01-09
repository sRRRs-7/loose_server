package session

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
)

var rClient *redis.Client

// auto execute before main function
func init() {
	rClient = redis.NewClient(&redis.Options{
		// Addr: "redis:6379",
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func NewRedis() *redis.Client {
	return rClient
}

// use login cookie hashToken == redis key
func NewSession(gc *gin.Context, cookieKey string, accessToken string, info []byte, redisExpired time.Duration, cookieExpired int) error {
	// hashToken, err := cryptography.HashPassword(accessToken)
	// if err != nil {
	// 	return fmt.Errorf("new session error cannot convert hash token: %v", err)
	// }
	// set redis
	rClient.Set(gc, accessToken, info, redisExpired)
	// set cookie
	// gc.SetCookie(cookieKey, accessToken, cookieExpired, "/", "localhost", true, true)
	gc.SetCookie(cookieKey, accessToken, cookieExpired, "/", "command-style.com", true, true)

	return nil
}

// check per access
func GetSession(gc *gin.Context, accessToken string) interface{} {
	// get redis value
	// hashToken, err := cryptography.HashPassword(accessToken)
	// if err != nil {
	// 	fmt.Printf("redis access token hash error : %v \n", err)
	// 	return nil
	// }

	redisValue, err := rClient.Get(gc, accessToken).Result()
	if err == redis.Nil {
		fmt.Printf("Get session error cannot get redis value because of value is not set: %v \n", err)
		return nil
	}
	if err != nil {
		fmt.Printf("Get session error cannot get redis value: %v \n", err)
		return nil
	}
	return redisValue
}

// usa logout
func DeleteSession(gc *gin.Context, cookieKey string) {
	redisId, _ := gc.Cookie(cookieKey)
	// delete redis record
	err := rClient.Del(gc, redisId)
	if err != nil {
		fmt.Printf("Delete session error cookie and redis cannot delete record")
	}
	// delete cookie
	// gc.SetCookie(cookieKey, "", -1, "/", "localhost", true, true)
	gc.SetCookie(cookieKey, "", -1, "/", "command-style.com", true, true)
}

// redis value get
func GetRedis(gc *gin.Context, key string) *redis.StringCmd {
	return rClient.Get(gc, key)
}
