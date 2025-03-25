package api

import (
	"net/http"

	"github.com/kirebyte/thd-project/internal/api/handlers"
	"github.com/kirebyte/thd-project/service"
)

// NewRouter creates and returns a configured HTTP router.
func NewRouter(svc service.Car) *http.ServeMux {
	mux := http.NewServeMux()
	carHandler := handlers.NewCarHandler(svc)
	carHandler.RegisterRoutes(mux)
	return mux
}
