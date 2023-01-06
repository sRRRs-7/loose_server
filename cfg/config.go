package cfg

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application
// The values are read by viper from config file or environment variables
type Config struct {
	DBdriver             string        `mapstructure:"DB_DRIVER"`
	DBsource             string        `mapstructure:"DB_SOURCE"`
	RedisSource          string        `mapstructure:"REDIS_SOURCE"`
	HttpServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GrpcServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	GinContextKey        string        `mapstructure:"GIN_CONTEXT_KEY"`
	RedisCookieKey       string        `mapstructure:"REDIS_COOKIE_KEY"`
	AdminCookieKey       string        `mapstructure:"REDIS_COOKIE_ADMIN_KEY"`
	AccessCookieDuration int           `mapstructure:"ACCESS_COOKIE_DURATION"`
	AccessRedisDuration  time.Duration `mapstructure:"ACCESS_REDIS_DURATION"`
}

// LoadConfig reads the configuration file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
