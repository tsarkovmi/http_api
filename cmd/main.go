package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //решает ошибку error init db: sql: unknown driver "postgres" (forgotten import?) exit status 1
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	httpapi "github.com/tsarkovmi/http_api"
	"github.com/tsarkovmi/http_api/pkg/handler"
	"github.com/tsarkovmi/http_api/pkg/repository"
	"github.com/tsarkovmi/http_api/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error init configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loadint env file: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("error init db: %s", err.Error())
	}

	//внедряем зависимости по порядку

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.Newhandler(services)

	srv := new(httpapi.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRourers()); err != nil {
		logrus.Fatalf("error running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
