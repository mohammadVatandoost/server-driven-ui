package main

import (
	"context"
	"sync"
	"syscall"
	"time"

	restAPI "github.com/mohammadVatandoost/server-driven-ui/internal/core/rest"
	cntext "github.com/mohammadVatandoost/server-driven-ui/pkg/context"
	"github.com/mohammadVatandoost/server-driven-ui/pkg/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start server",
	Run: func(cmd *cobra.Command, args []string) {
		if err := serve(cmd, args); err != nil {
			logrus.WithError(err).Fatal("Failed to serve.")
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) error {
	conf := loadConfigOrPanic(cmd)
	configureLoggerOrPanic(conf.Logger)

	log := logger.NewLogger()

	serverREST := restAPI.New(log, conf.Rest)
	serverREST.Routes()

	serverContext, serverCancel := cntext.WithSignalCancellation(
		context.Background(),
		syscall.SIGTERM, syscall.SIGINT,
	)
	defer serverCancel()

	var serverWaitGroup sync.WaitGroup
	serverWaitGroup.Add(1)

	go func() {
		defer serverWaitGroup.Done()
		serverREST.Run()
	}()
	<-serverContext.Done()
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		serverREST.Shutdown(ctx)
	}()

	serverWaitGroup.Wait()
	return nil
}
