package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/kirebyte/thd-project/internal/logger"
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
func (s *Car) Get(ctx context.Context, id string) (model.Car, error) {
	// Validate the car ID.
	if id == "" {
		return model.Car{}, errors.New("car ID must not be empty")
	}

	// Retrieve the car from the repository.
	car, err := s.repo.FindByID(ctx, id)
	if err != nil {
		logger.Error(fmt.Sprintf("error while getting car: %v", err))
		return model.Car{}, errors.New("error while getting, please contact the administrator")
	}

	return car, nil
}

// List returns all cars currently stored.
func (s *Car) List(ctx context.Context) ([]model.Car, error) {
	// Retrieve all cars from the repository.
	cars, err := s.repo.FindAll(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error while listing cars: %v", err))
		return nil, errors.New("error while listing, please contact the administrator")
	}
	return cars, nil
}

// Create saves a new car to the repository.
func (s *Car) Create(ctx context.Context, car model.Car) (model.Car, error) {
	// Validate the car's fields.
	if err := validateCarFields(car); err != nil {
		return model.Car{}, err
	}

	// Generate a unique ID for the car based on the current timestamp.
	timestamp := time.Now().UnixNano()
	hash := sha256.Sum256([]byte(strconv.FormatInt(timestamp, 10)))
	car.ID = hex.EncodeToString(hash[:])

	err := s.repo.Save(ctx, car)
	if err != nil {
		logger.Error(fmt.Sprintf("error while saving car: %v", err))
		return model.Car{}, errors.New("error while saving, please contact the administrator")
	}

	// Save the car to the repository.
	return car, nil
}

// Update updates an existing car by ID.
func (s *Car) Update(ctx context.Context, car model.Car) error {
	// Validate the car's ID.'
	if car.ID == "" {
		return errors.New("missing ID")
	}

	// Validate the car's fields.
	if err := validateCarFields(car); err != nil {
		return err
	}

	// Update the car in the repository.
	err := s.repo.Update(ctx, car)
	if err != nil {
		logger.Error(fmt.Sprintf("error while updating car: %v", err))
		return errors.New("error while updating, please contact the administrator")
	}
	return nil
}
