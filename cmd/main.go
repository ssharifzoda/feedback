package main

import (
	"feedback/internal/api/handlers"
	"feedback/internal/botSystem"
	"feedback/internal/database"
	"feedback/internal/server"
	"feedback/internal/service"
	"feedback/pkg/logging"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	logger := logging.GetLogger()
	if err := initConfig(); err != nil {
		logger.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error initializing env value: %s", err.Error())
	}
	conn, err := database.NewPostgresGorm()
	if err != nil {
		logger.Fatalf("failed to initializing db: %s", err.Error())
	}
	repository := database.NewDatabase(conn)
	services := service.NewService(repository, viper.GetString("storage.imagepath"))
	go botSystem.NewBotSystem(services)
	newHandler := handlers.NewHandler(services)
	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), newHandler.InitRoutes()); err != nil {
		logger.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("internal/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
