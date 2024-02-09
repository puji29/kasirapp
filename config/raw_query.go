package config

const (
	//product
	InsertProduct = `INSERT INTO product (name, harga, stock) VALUES (?, ?, ?) Retruning id, creadted_at`
)
