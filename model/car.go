package model

type Car struct {
	ID       string `json:"id"`
	Make     string `json:"make"`
	Model    string `json:"model"`
	Package  string `json:"package"`
	Color    string `json:"color"`
	Year     int    `json:"year"`
	Category string `json:"category"`
	Mileage  int    `json:"mileage"` // miles
	Price    int    `json:"price"`   // cents
}
