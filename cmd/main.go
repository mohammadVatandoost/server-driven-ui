package main

import (
	"fmt"
	"os"

	"github.com/mohammadVatandoost/server-driven-ui/internal/config"
	"github.com/mohammadVatandoost/server-driven-ui/pkg/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const serviceName = "server_driven_ui"

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func loadConfigOrPanic(cmd *cobra.Command) *config.Config {
	conf, err := config.LoadConfig(cmd)
	if err != nil {
		logrus.WithError(err).Panic("Failed to load configurations")
	}
	return conf
}

func configureLoggerOrPanic(loggerConfig logger.Config) {
	if err := logger.Initialize(&loggerConfig); err != nil {
		logrus.WithError(err).Panic("Failed to configure logger")
	}
}
