package main

import (
	"effectiveMobile"
	"effectiveMobile/configs"
	"effectiveMobile/pkg/database"
	"effectiveMobile/pkg/httpserver/handlers"
	"effectiveMobile/pkg/util/log/sl"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := setLog()
	config := configs.LoadConfigs(log)

	db, err := database.NewDatabase(database.DatabaseLoad{
		Host:     config.DB.Host,
		Port:     config.DB.Port,
		DBName:   config.DB.Name,
		User:     config.DB.Username,
		Password: config.DB.Password,
	})
	if err != nil {
		log.Error("could not init db", sl.Err(err))
		os.Exit(1)
	}

	handlerBody := handlers.HandlerInitBody{
		Log: log,
		DB:  db,
		Urls: handlers.HandlerInitUrls{
			AgeUrl:         config.AgeHost,
			GenderUrl:      config.GenderHost,
			NationalityUrl: config.NationalityHost,
		},
	}

	server := new(effectiveMobile.Server)
	go func() {
		if err := server.Run(config.ServerPort, handlers.InitRoutes(handlerBody)); err != nil {
			log.Error("could not start server", sl.Err(err))
			os.Exit(1)
		}
	}()

	log.Info("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := server.Close(); err != nil {
		log.Error("could not shutdown server", sl.Err(err))
		os.Exit(1)
	}

	log.Info("server stopped")
}

func setLog() *slog.Logger {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	return log
}
