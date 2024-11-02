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
	/*
		Включаем логирование, и считываем конфиг
		с помощью библиотеки viper
	*/
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.Info("Initializing configuration with Viper")
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init configs: %s", err.Error())
	}
	/*
		Считываем .env файл
	*/
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loadint env file: %s", err.Error())
	}
	/*
		Здесь инициализируем репозиторий, считываем во все необходимые поля
		То, что считали выше, теперь записываем в поля структуры Config
	*/
	logrus.Info("Connecting to the database")
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

	/*
		По порядку поднимаем сервис
		Сначала передаем указатель на БД метод NewRepository
		Который просто оборачивает БД в структуру Repostory, чтобы связать все слои интерфейсами
		Аналогичная процедура делается для Сервиса и Хэндлера
		Таким образом все слои связаны и общаются друг с другом с помощью интерфейсов
	*/
	logrus.Info("Initializing repository")
	repos := repository.NewRepository(db)
	logrus.Info("Initializing service layer")
	services := service.NewService(repos)
	logrus.Info("Initializing handlers")
	handlers := handler.Newhandler(services)

	/*
		Инициализируем сервер
		А далее запускаем с помощью метода Run
		В методе Run передаём порт, на котором будет запускаться сервер
		А также передаем Хэндлер, сразу же инициализируя его создавая необходимые HTTP запросы
		Возвращаем метод ListenAndServe
		Логируем в случае ошибки
	*/
	logrus.Info("Starting HTTP server")
	srv := new(httpapi.Server)
	//GRACEFUL SHUTDOWN
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRourers()); err != nil {
			logrus.Fatalf("error running http server: %s", err.Error())
		}
	}()
	logrus.Print("http app Started")
	/*
		Создаем GRACEFUL SHUTDOWN
		Канал quit, который будет удерживать main от завершения работы
		Как только будет получен SIGTERM или SIGINT канал закроется и мейн пойдёт дальше
		Логируем завершение программы
	*/
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("http app Shutting Down")
	/*
		Теперь вызывается метод Shutdown, который поочереди закрывает все листенеры, не прерывая активных подключений,
		затем закрывает все бездействующие подключения. Ожидает пока все подключения не вернутся в состояние бездействия,
		а затем завершает работу
		Далее закрывает подключение к БД.
	*/
	//То, что возвращается Fatal - нормально
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

/*
Считываем из Директории configs
файл с названием config
*/
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
