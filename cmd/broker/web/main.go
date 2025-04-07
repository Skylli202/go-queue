package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Skylli202/go-queue/broker"
	"github.com/Skylli202/go-queue/broker/store"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(
	ctx context.Context,
	args []string,
	getenv func(string) string,
	stdin io.Reader,
	stdout, stderr io.Writer,
) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// TODO: Add control of log level from both ENV & ARGS
	opts := &slog.HandlerOptions{Level: slog.LevelDebug}
	handler := slog.NewJSONHandler(stdout, opts)
	slogger := slog.New(handler)

	// TODO: Add DB path control from both ENV & ARGS
	db, err := sql.Open("sqlite3", "./main.db")
	if err != nil {
		return err
	}
	defer db.Close()
	stores := store.NewStores(db)

	s := broker.NewBrokerServer(
		&broker.BrokerServerConfig{},
		slogger,
		stores,
	)
	// TODO: Add server's port control from both ENV & ARGS
	httpServer := &http.Server{
		Addr:    net.JoinHostPort("", "8080"),
		Handler: s,
	}

	go func() {
		fmt.Fprintf(stdout, "Server listening on address: %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(stderr, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()

	return nil
}
