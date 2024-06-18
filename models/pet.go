package models

import "time"

type Status string
type Category string

const (
	Available Status = "available"
	Pending Status = "pending"
	Sold Status = "sold"

	Dog Category = "dog"
	Cat Category = "cat"
	Bird Category = "bird"
	Fish Category = "fish"
	Reptile Category = "reptile"
	Other Category = "other"
)

type Pet struct {
	ID int64 `json:"id"`
	Name string `json:"name" validate:"required"`
	Status string `json:"status" validate:"required,status"`
	Category string `json:"category" validate:"required,category"`
	Tags []string `json:"tags,omitempty"`
	Age int	`json:"age,omitempty"`
	Color string `json:"color,omitempty"`
	Gender string `json:"gender,omitempty"`
	Weight float64 `json:"weight,omitempty"`
	DateAdded time.Time `json:"dateAdded,omitempty"`
}