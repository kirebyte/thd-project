package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/kirebyte/thd-project/internal/logger"
	"github.com/kirebyte/thd-project/model"
	"github.com/kirebyte/thd-project/service"
)

type CarHandler struct {
	svc service.Car
}

// NewCarHandler creates a new handler with the given CarService.
func NewCarHandler(svc service.Car) *CarHandler {
	return &CarHandler{svc: svc}
}

// RegisterRoutes adds the car endpoints to the provided mux.
func (h *CarHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /cars/", h.GetCar)
	mux.HandleFunc("GET /cars", h.ListCars)
	mux.HandleFunc("POST /cars", h.CreateCar)
	mux.HandleFunc("PUT /cars/", h.UpdateCar)
}

// GET /cars/{id}
func (h *CarHandler) GetCar(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/cars/")
	if id == "" {
		http.Error(w, "Missing car ID", http.StatusBadRequest)
		return
	}

	car, err := h.svc.Get(r.Context(), id)
	if err != nil {
		logger.Warn("GetCar failed: " + err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	writeJSON(w, http.StatusOK, car)
}

// GET /cars
func (h *CarHandler) ListCars(w http.ResponseWriter, r *http.Request) {
	cars, err := h.svc.List(r.Context())
	if err != nil {
		logger.Error("ListCars failed: " + err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, cars)
}

// POST /cars
func (h *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a car.
	var car model.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Create the car.
	addedCar, err := h.svc.Create(r.Context(), car)
	if err != nil {
		logger.Warn("CreateCar failed: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, addedCar)
}

// PUT /cars/{id}
func (h *CarHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	var car model.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.svc.Update(r.Context(), car); err != nil {
		logger.Warn("UpdateCar failed: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusOK, car)
}

// writeJSON serializes the response as JSON.
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
