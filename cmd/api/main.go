package main

import (
	"net"

	"github.com/OlenEnkeli/GoCurrency/internal/api"
	"github.com/OlenEnkeli/GoCurrency/internal/settings"
	"github.com/OlenEnkeli/GoCurrency/internal/utils"
	"github.com/sirupsen/logrus"
)

func main() {
	utils.SetupProject()

	server := api.NewServer()

	apiUrl := net.JoinHostPort(
		settings.Settings.API.Host,
		settings.Settings.API.Port,
	)
	logrus.Infof("Running API at %s, [%s mode]", apiUrl, settings.Settings.App.Mode)
	logrus.Infof("Swagger awailable at %s/swagger/index.html", apiUrl)

	if err := server.Run(); err != nil {
		logrus.Fatalf("Fatal error: %s", err)
	}
}
