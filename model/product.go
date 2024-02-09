package model

import "time"

type Product struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Harga      int       `json:"harga"`
	Stock      int       `json:"stock"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
