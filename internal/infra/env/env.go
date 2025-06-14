package env

import (
	"github.com/spf13/viper"
)

type Env struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
		SSLMode  string `json:"sslmode"`
	} `json:"database"`

	App struct {
		Port         string `json:"port"`
		Version      string `json:"version"`
		Name         string `json:"name"`
		Env          string `json:"env"`
		ApiKey       string `json:"apiKey"`
		ApiKeyHeader string `json:"apiKeyHeader"`
	} `json:"app"`

	RabbitMQ struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"rabbitmq"`

	Logger struct {
		Type     string `json:"type"`
		WithFile bool   `json:"withFile"`
	} `json:"logger"`

	Jwt struct {
		ATSecretKey   string `json:"atSecretKey"`
		ATExpiredTime int    `json:"atExpiredTime"`
		RTSecretKey   string `json:"rtSecretKey"`
		RTExpiredTime int    `json:"rtExpiredTime"`
	} `json:"jwt"`
}

func NewEnv() *Env {
	env := Env{}

	viper.SetConfigFile("env.json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		panic(err)
	}

	return &env
}

func NewEnvFromFile(filepath string) *Env {
	env := Env{}

	viper.SetConfigFile(filepath)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		panic(err)
	}

	return &env
}
