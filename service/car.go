package service

import (
	"context"

	"github.com/kirebyte/thd-project/model"
)

type Car interface {
	Get(ctx context.Context, id string) (model.Car, error)
	List(ctx context.Context) ([]model.Car, error)
	Create(ctx context.Context, car model.Car) (model.Car, error)
	Update(ctx context.Context, car model.Car) error
}
