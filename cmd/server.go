package main

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/canpacis/go-assignment/internal/errs"
	"github.com/canpacis/go-assignment/services/greet"
)

func main() {
	mux := http.NewServeMux()

	greetService := &greet.Service{}

	// Inlining this http.Handler so as not to overengineer a solution.
	// A helper function that can generate & validate the payload, run the
	// service method and write any result/error to the client can replace this.
	mux.HandleFunc("GET /hello-world", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		query := r.URL.Query()

		resp, err := greetService.Greet(ctx, &greet.GreetPayload{Name: query.Get("name")})
		w.Header().Set("Content-Type", "application/json")

		encoder := json.NewEncoder(w)

		if err != nil {
			var apiErr *errs.Error
			if errors.As(err, &apiErr) {
				w.WriteHeader(apiErr.Status)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}

			encoder.Encode(err)
		} else {
			w.WriteHeader(http.StatusOK)
			encoder.Encode(resp)
		}
	})

	Serve(mux)
}

// Handles running and the graceful shutdown of the server
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
