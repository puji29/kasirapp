package config

const (
	//product
	InsertProduct = `INSERT INTO produk (nama, harga, stock) VALUES ($1, $2, $3) RETURNING id, created_at;`
	SelectProduct = `SELECT id,nama, harga,stock,created_at, update_at FROM produk`
)
