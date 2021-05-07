package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

var Environment environment

type environment struct {
	Host         string
	Port         string
	Env          string
	Url          string
	CacheViews   string
	Key          string
	Iv           string
	DbConnection string `split_words:"true"`
	DbHost       string `split_words:"true"`
	DbPort       string `split_words:"true"`
	DbUser       string `split_words:"true"`
	DbPassword   string `split_words:"true"`
	DbName       string `split_words:"true"`
	DbSslMode    string `split_words:"true"`
}

func InitEnvironment() {
	if err := godotenv.Load("pkg/config/.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := envconfig.Process("app", &Environment); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(Environment)
}
