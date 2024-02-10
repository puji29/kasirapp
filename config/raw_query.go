package config

const (
	//product
	InsertProduct = `INSERT INTO produk (nama, harga, stock) VALUES ($1, $2, $3) RETURNING id, created_at;`
)
