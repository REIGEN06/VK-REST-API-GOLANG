package main

import (
	"context"
	"github.com/reigen06/vk-rest-api/config"
	"github.com/reigen06/vk-rest-api/internal/common/logger"
	"github.com/reigen06/vk-rest-api/internal/handler"
	"github.com/reigen06/vk-rest-api/internal/models"
	"github.com/reigen06/vk-rest-api/internal/repository"
	"github.com/reigen06/vk-rest-api/internal/service"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

// @title VK-REST-API API
// @version 1.0
// @description API Server for VK-REST-API Application

// @host localhost:4005
// @BasePath /
func main() {
	cfg := config.NewConfig()

	log := logger.SetupLogger(cfg.Env)

	log.Info("starting vk-rest-api...", slog.String("env", cfg.Env))

	db, err := repository.NewPostgresDB(config.DBConfig{
		Host:     cfg.DBConfig.Host,
		Port:     cfg.DBConfig.Port,
		Username: cfg.DBConfig.Username,
		Database: cfg.DBConfig.Database,
		Password: cfg.DBConfig.Password,
	})
	if err != nil {
		log.Error("failed to initialize db", slog.String("cause", err.Error()))
		os.Exit(1)
	}

	if err := db.AutoMigrate(&models.Actor{}); err != nil {
		log.Error("failed to migrate", slog.String("cause", err.Error()))
		os.Exit(1)
	}

	repos := repository.NewRepository(db, log)
	services := service.NewService(repos, log)
	handlers := handler.NewHandler(services, log)

	srv := new(config.Server)
	go func() {
		if err := srv.Run(handlers.SetRoutes(), *cfg); err != nil {
			log.Error("error occurred while running http server:", slog.String("cause", err.Error()))
			os.Exit(1)
		}
	}()

	log.Info("Server listening at ", slog.String("address", "http://"+cfg.Server.HttpServer.Addr))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info("Server Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Error("error occurred on server shutting down:", slog.String("cause", err.Error()))
	}

	sqlDB, err := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Error("error occurred on db connection close:", slog.String("cause", err.Error()))
	}

}
