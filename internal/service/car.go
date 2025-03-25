package service

import (
	"context"
	"errors"

	"github.com/kirebyte/thd-project/model"
	"github.com/kirebyte/thd-project/repository"
)

// Car provides business logic around car operations.
type Car struct {
	repo repository.Car
}

// New creates a new CarService with the given repository.
func New(repo repository.Car) *Car {
	return &Car{repo: repo}
}

// Get retrieves a car by its ID.
//
// Precondition: id must not be empty.
// Postcondition: returns the matching car or an error if not found.
func (s *Car) Get(ctx context.Context, id string) (model.Car, error) {
	if id == "" {
		return model.Car{}, errors.New("car ID must not be empty")
	}
	return s.repo.FindByID(ctx, id)
}

// List returns all cars currently stored.
func (s *Car) List(ctx context.Context) ([]model.Car, error) {
	return s.repo.FindAll(ctx)
}

// Create saves a new car to the repository.
//
// Precondition: car.ID must not be empty.
// Postcondition: car is persisted in the repository.
func (s *Car) Create(ctx context.Context, car model.Car) error {
	if car.ID == "" {
		return errors.New("car ID is required")
	}
	// Optional: add more validation here (e.g. year > 1900)
	return s.repo.Save(ctx, car)
}

// Update updates an existing car by ID.
//
// Precondition: id and car.ID must match and be non-empty.
// Postcondition: car data is updated in the repository.
func (s *Car) Update(ctx context.Context, id string, car model.Car) error {
	if id == "" || car.ID == "" {
		return errors.New("both ID and car.ID must be provided")
	}
	if id != car.ID {
		return errors.New("path ID and car.ID must match")
	}
	return s.repo.Update(ctx, id, car)
}
