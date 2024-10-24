package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //решает ошибку error init db: sql: unknown driver "postgres" (forgotten import?) exit status 1
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	httpapi "github.com/tsarkovmi/http_api"
	_ "github.com/tsarkovmi/http_api/docs"
	"github.com/tsarkovmi/http_api/pkg/handler"
	"github.com/tsarkovmi/http_api/pkg/repository"
	"github.com/tsarkovmi/http_api/pkg/service"
)

//	@title			Workers Management API
//	@version		1.0
//	@description	API для управления данными о работниках (CRUD операции).
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/
//	@schemes	http

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
	//GRACEFUL SHUTDOWN
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRourers()); err != nil {
			logrus.Fatalf("error running http server: %s", err.Error())
		}
	}()
	logrus.Print("http app Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("http app Shutting Down")

	//То, что возвращается Fatal - нормально, так и должно быть
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
