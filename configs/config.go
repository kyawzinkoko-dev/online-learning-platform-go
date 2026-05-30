package configs

import "github.com/spf13/viper"

type Config struct {
	AppPort     string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBSSLMode   string
	Environment string

	JWTSecret string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	config := &Config{
		AppPort: viper.GetString("APP_PORT"),

		DBHost:      viper.GetString("DB_HOST"),
		DBPort:      viper.GetString("DB_PORT"),
		DBUser:      viper.GetString("DB_USER"),
		DBPassword:  viper.GetString("DB_PASSWORD"),
		DBName:      viper.GetString("DB_NAME"),
		DBSSLMode:   viper.GetString("DB_SSLMODE"),
		Environment: viper.GetString("ENVIRONMENT"),
		JWTSecret:   viper.GetString("JWT_SECRET"),
	}
	return config, nil
}
