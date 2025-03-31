package env

import "github.com/spf13/viper"

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
		Port    string `json:"port"`
		Version string `json:"version"`
		Name    string `json:"name"`
	} `json:"app"`

	Logger struct {
		Type string `json:"type"`
	}

	Jwt struct {
		ATSecretKey string `json:"accessTokenSecretKey"`
		ATExpiredTime int `json:"accessTokenExpiredTime"`
		RTSecretKey string `json:"refreshTokenSecretKey"`
		RTExpiredTime int `json:"refreshTokenExpiredTime"`
	} `json:"jwt"`
}

func NewEnv() *Env {
	env := Env{}

	viper.SetConfigFile("./env.json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		panic(err)
	}

	return &env
}