package main

import (
	"context"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"yaus"
	"yaus/internal/handler"
	"yaus/internal/repository"
	"yaus/internal/service"
)

// @title 			Yet Another URL Shortener (yaus)
// @version  		0.1.0
// @description     Сервис сокращения URL'ов
// @host      		0.0.0.0:8080
// @BasePath  		/api
func main() {
	rand.Seed(time.Now().UnixNano())
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := loadEnv(); err != nil {
		logrus.Fatalf("error while load .env: %s", err.Error())
	}

	db, err := initDB()
	if err != nil {
		logrus.Fatalf("error while initializing db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(yaus.Server)
	go func() {
		if err := server.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != http.ErrServerClosed {
			logrus.Fatalf("error, while starting http server: %s", err.Error())
		}
	}()

	logrus.Println("Server started work")

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	logrus.Println("Server shutting down")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error, while shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error, while closing connecton to DB: %s", err.Error())
	}
}

func loadEnv() error {
	err := godotenv.Load()
	return err
}

func initDB() (db *sqlx.DB, err error) {
	db, err = repository.NewPostgresDB(repository.Config{
		Host:    os.Getenv("DB_HOST"),
		Port:    os.Getenv("DB_PORT"),
		User:    os.Getenv("DB_USER"),
		Pass:    os.Getenv("DB_PASSWORD"),
		NameDB:  os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSL"),
	})
	return
}
