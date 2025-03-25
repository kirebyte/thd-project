package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/kirebyte/thd-project/internal/model"
)

type CarRepository struct {
	db *sql.DB
}

// New creates a new CarRepository
func New(db *sql.DB) *CarRepository {
	return &CarRepository{db: db}
}

// FindByID returns a car by its ID
func (r *CarRepository) FindByID(ctx context.Context, id string) (model.Car, error) {
	row := r.db.QueryRowContext(ctx, `
		SELECT id, make, model, package, color, year, category, mileage, price
		FROM cars WHERE id = ?`, id)

	var car model.Car
	err := row.Scan(&car.ID, &car.Make, &car.Model, &car.Package, &car.Color, &car.Year, &car.Category, &car.Mileage, &car.Price)
	if err != nil {
		return model.Car{}, fmt.Errorf("FindByID error: %w", err)
	}
	return car, nil
}

// FindAll returns all cars
func (r *CarRepository) FindAll(ctx context.Context) ([]model.Car, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, make, model, package, color, year, category, mileage, price
		FROM cars`)
	if err != nil {
		return nil, fmt.Errorf("FindAll error: %w", err)
	}
	defer rows.Close()

	var cars []model.Car
	for rows.Next() {
		var car model.Car
		if err := rows.Scan(&car.ID, &car.Make, &car.Model, &car.Package, &car.Color, &car.Year, &car.Category, &car.Mileage, &car.Price); err != nil {
			return nil, fmt.Errorf("row scan error: %w", err)
		}
		cars = append(cars, car)
	}
	return cars, nil
}

// Save inserts a new car
func (r *CarRepository) Save(ctx context.Context, car model.Car) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO cars (id, make, model, package, color, year, category, mileage, price)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		car.ID, car.Make, car.Model, car.Package, car.Color, car.Year, car.Category, car.Mileage, car.Price)
	if err != nil {
		return fmt.Errorf("Save error: %w", err)
	}
	return nil
}

// Update updates a car
func (r *CarRepository) Update(ctx context.Context, id string, car model.Car) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE cars
		SET make = ?, model = ?, package = ?, color = ?, year = ?, category = ?, mileage = ?, price = ?
		WHERE id = ?`,
		car.Make, car.Model, car.Package, car.Color, car.Year, car.Category, car.Mileage, car.Price, id)
	if err != nil {
		return fmt.Errorf("Update error: %w", err)
	}
	return nil
}
