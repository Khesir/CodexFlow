package main

import (
	"app/internal/migrations"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
)

// MigrationFunc defines the function signature for migrations
type MigrationFunc func(db *sqlx.DB) error

// Migration represents a single migration
type Migration struct {
	Version string `json:"version"`
}
type App struct {
	DB         *sqlx.DB
	Migrations map[string]MigrationFunc
}

// return initial configs
// e.g. constant configs
// url config
// i18n
// Media Storage
// core
// auth manager
// other necessary features that doesnt involved the api
func main() {
	db := initDB("localhost", 5432, "postgres", "password", "mydb")

	migrations.RunMigrations(db)

	app := &App{
		DB: db,
	}
	r := initHTTPServer(app)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("Server started on :8080")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	db.Close()
	log.Println("Server exiting")
}
