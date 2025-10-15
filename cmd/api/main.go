package main

import (
	"net"

	"github.com/OlenEnkeli/GoCurrency/internal/controllers"
	"github.com/OlenEnkeli/GoCurrency/internal/handlers/api"
	"github.com/OlenEnkeli/GoCurrency/internal/repositories"
	"github.com/OlenEnkeli/GoCurrency/internal/settings"
	"github.com/OlenEnkeli/GoCurrency/internal/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	utils.SetupSettingsAndLogger()

	db := repositories.NewPostgresConnection()
	repository := repositories.NewRepository(db)
	contoller := controllers.NewController(repository)

	apiHandler := api.NewHandler(contoller)
	apiServer := apiHandler.NewServer()

	apiUrl := net.JoinHostPort(
		settings.Settings.API.Host,
		settings.Settings.API.Port,
	)
	logrus.Infof("Running API at %s, [%s mode]", apiUrl, settings.Settings.App.Mode)
	logrus.Infof("Swagger awailable at %s/swagger/index.html", apiUrl)

	if err := apiServer.Run(); err != nil {
		logrus.Fatalf("Fatal error: %s", err)
	}
}
