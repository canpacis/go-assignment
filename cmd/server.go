package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()

	Serve(mux)
}

func Serve(h http.Handler) {
	server := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Server listen failed", "error", err)
			os.Exit(1)
		}
	}()
	slog.Info("Server is running", "address", server.Addr)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	ctx, release := context.WithTimeout(context.Background(), time.Second*10)
	defer release()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown failed", "error", err)
		os.Exit(1)
	}
	slog.Info("Server shutdown")
}
