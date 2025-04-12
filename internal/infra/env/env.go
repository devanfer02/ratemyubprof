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

	Logger struct {
		Type string `json:"type"`
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
