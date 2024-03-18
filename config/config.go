package config

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

type Config struct {
	Env string
	DBConfig
	Server
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type Server struct {
	HttpServer *http.Server
}

func NewConfig() *Config {
	envLoaded := true

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
		envLoaded = false
	}

	if !envLoaded {
		return &Config{
			Env: "dev",
			DBConfig: DBConfig{
				"postgres-database",
				"5432",
				"postgres",
				"password",
				"film-library",
			},
			Server: Server{
				&http.Server{
					Addr:           "localhost:4005",
					Handler:        nil,
					MaxHeaderBytes: 1 << 20,
					ReadTimeout:    10 * time.Second,
					WriteTimeout:   10 * time.Second},
			},
		}
	}

	return &Config{
		Env: os.Getenv("ENV"),
		DBConfig: DBConfig{
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
		},
		Server: Server{
			&http.Server{
				Addr:           os.Getenv("API_HOST") + ":" + os.Getenv("PORT"),
				Handler:        nil,
				MaxHeaderBytes: 1 << 20,
				ReadTimeout:    10 * time.Second,
				WriteTimeout:   10 * time.Second},
		},
	}
}

func (s *Server) Run(handler http.Handler, cfg Config) error {
	s.HttpServer = &http.Server{
		Addr:           cfg.Server.HttpServer.Addr,
		Handler:        handler,
		MaxHeaderBytes: cfg.Server.HttpServer.MaxHeaderBytes,
		ReadTimeout:    cfg.Server.HttpServer.ReadTimeout,
		WriteTimeout:   cfg.Server.HttpServer.WriteTimeout,
	}

	return s.HttpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.HttpServer.Shutdown(ctx)
}
