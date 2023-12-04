package main

import (
	"os"

	auth "github.com/eeQuillibrium/go-rest-auth"
	"github.com/eeQuillibrium/go-rest-auth/pkg/handler"
	"github.com/eeQuillibrium/go-rest-auth/pkg/repository"
	"github.com/eeQuillibrium/go-rest-auth/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatal().
		AnErr("Config initializing error", err)
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal().
		AnErr("Env load error", err)
	}

	db, err := repository.StartPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatal().
			Err(err).AnErr("db pinging problem", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	fserver := new(auth.FileServer)
	serv := new(auth.Server)

	if err := serv.Run(viper.GetString("port"), handlers.InitRoutes(fserver.Initialize())); err != nil {
		log.Fatal().
			Err(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
