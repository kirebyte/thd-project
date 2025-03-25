package repository

import (
	"context"

	"github.com/kirebyte/thd-project/model"
)

type Car interface {
	FindByID(ctx context.Context, id string) (model.Car, error)
	FindAll(ctx context.Context) ([]model.Car, error)
	Save(ctx context.Context, car model.Car) error
	Update(ctx context.Context, id string, car model.Car) error
}
