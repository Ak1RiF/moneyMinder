package app

import (
	"MoneyMinder"
	"MoneyMinder/internal/handler"
	"MoneyMinder/internal/helpers"
	"MoneyMinder/internal/repository"
	"MoneyMinder/internal/service"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	if err := helpers.InitConfig(); err != nil {
		log.Fatal().Msgf("Error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DB:       viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatal().Msgf("Failed connection to db: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	server := new(MoneyMinder.Server)
	go func() {
		if err := server.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
			log.Fatal().Msgf("Failed attempt to start the server: %s", err.Error())
		}
	}()

	log.Info().Msg("MoneyMinder server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
