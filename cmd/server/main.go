package main

import (
	"database/sql"
	"net/http"

	"github.com/kirebyte/thd-project/internal/api"
	"github.com/kirebyte/thd-project/internal/logger"
	"github.com/kirebyte/thd-project/internal/repository/sqlite"
	"github.com/kirebyte/thd-project/internal/service"
	"github.com/kirebyte/thd-project/settings"
	_ "modernc.org/sqlite"
)

func main() {
	cfg := settings.Load()

	// Open database connection
	db, err := sql.Open("sqlite", cfg.DBPath)
	if err != nil {
		logger.Fatal("Failed to connect to DB: " + err.Error())
	}
	defer db.Close()

	// Init repository
	repo := sqlite.New(db)

	// Init service
	carService := service.New(repo)

	// Init router
	router := api.NewRouter(carService)

	// Start server
	addr := ":" + cfg.Port
	logger.Info("🚗 Starting server on " + addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		logger.Fatal("Server failed: " + err.Error())
	}
}
