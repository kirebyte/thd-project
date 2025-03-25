package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kirebyte/thd-project/internal/api/handlers"
	"github.com/kirebyte/thd-project/model"
)

// mockCarService implements service.Car for testing purposes
type mockCarService struct {
	GetFunc    func(ctx context.Context, id string) (model.Car, error)
	ListFunc   func(ctx context.Context) ([]model.Car, error)
	CreateFunc func(ctx context.Context, car model.Car) (model.Car, error)
	UpdateFunc func(ctx context.Context, car model.Car) error
}

func (m *mockCarService) Get(ctx context.Context, id string) (model.Car, error) {
	return m.GetFunc(ctx, id)
}
func (m *mockCarService) List(ctx context.Context) ([]model.Car, error) {
	return m.ListFunc(ctx)
}
func (m *mockCarService) Create(ctx context.Context, car model.Car) (model.Car, error) {
	return m.CreateFunc(ctx, car)
}
func (m *mockCarService) Update(ctx context.Context, car model.Car) error {
	return m.UpdateFunc(ctx, car)
}

func TestGetCar_OK(t *testing.T) {
	mockSvc := &mockCarService{
		GetFunc: func(ctx context.Context, id string) (model.Car, error) {
			return model.Car{ID: id, Make: "Toyota", Model: "Corolla", Year: 2020}, nil
		},
	}
	handler := handlers.NewCarHandler(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/cars/abc123", nil)
	rr := httptest.NewRecorder()

	handler.GetCar(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 OK, got %d", rr.Code)
	}
	var car model.Car
	if err := json.NewDecoder(rr.Body).Decode(&car); err != nil {
		t.Fatalf("error decoding response: %v", err)
	}
	if car.ID != "abc123" {
		t.Errorf("expected ID abc123, got %s", car.ID)
	}
}

func TestCreateCar_BadJSON(t *testing.T) {
	mockSvc := &mockCarService{}
	handler := handlers.NewCarHandler(mockSvc)

	req := httptest.NewRequest(http.MethodPost, "/cars", bytes.NewBufferString(`not json`))
	rr := httptest.NewRecorder()

	handler.CreateCar(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 BadRequest, got %d", rr.Code)
	}
}

func TestCreateCar_Valid(t *testing.T) {
	mockSvc := &mockCarService{
		CreateFunc: func(ctx context.Context, car model.Car) (model.Car, error) {
			if car.Make == "" {
				return model.Car{}, errors.New("missing make")
			}
			return model.Car{}, nil
		},
	}
	handler := handlers.NewCarHandler(mockSvc)

	payload := model.Car{
		ID:       "abc123",
		Make:     "Mazda",
		Model:    "3",
		Year:     2022,
		Category: "sedan",
		Package:  "sport",
		Color:    "red",
		Mileage:  5000,
		Price:    2200000,
	}
	data, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPost, "/cars", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.CreateCar(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected 201 Created, got %d", rr.Code)
	}
}
