package service

import (
	"errors"

	"github.com/kirebyte/thd-project/model"
)

// validateCarFields checks if all fields of a car are non-empty or non-zero.
func validateCarFields(car model.Car) error {
	if car.Make == "" ||
		car.Model == "" ||
		car.Package == "" ||
		car.Color == "" ||
		car.Category == "" ||
		car.Year == 0 ||
		car.Mileage == 0 ||
		car.Price == 0 {
		return errors.New("all car fields must be non-empty or non-zero")
	}

	if car.Year < 1900 {
		return errors.New("car year must be 1900 or later")
	}

	return nil
}
