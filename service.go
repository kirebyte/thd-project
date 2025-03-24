package main

import (
	"context"

	"github.com/kirebyte/thd-project/internal/model"
)

type CarService interface {
	GetCar(ctx context.Context, id string) (model.Car, error)
	ListCars(ctx context.Context) ([]model.Car, error)
	CreateCar(ctx context.Context, car model.Car) error
	UpdateCar(ctx context.Context, id string, car model.Car) error
}
