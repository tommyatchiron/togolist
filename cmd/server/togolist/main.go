package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tommyatchiron/togolist/internal/app"
)

func main() {

	// Create context that listens for the interrupt signal from the OS
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	mainApp := app.New()

	// Graceful shutdown
	// https://github.com/gin-gonic/gin#manually

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := mainApp.Start(ctx); err != nil {
			fmt.Printf("%s\t[FATAL]\tError in staring application.\n", time.Now().Format(time.RFC3339))
			os.Exit(1)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	fmt.Printf("%s\t[INFO]\tShutting down gracefully, press Ctrl+C again to force\n", time.Now().Format(time.RFC3339))

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := mainApp.Stop(ctx); err != nil {
		fmt.Printf("%s\t[FATAL]\tServer forced to shutdown.\n", time.Now().Format(time.RFC3339))
		os.Exit(1)
	}

	fmt.Printf("%s\t[INFO]\tServer exiting\n", time.Now().Format(time.RFC3339))
}
