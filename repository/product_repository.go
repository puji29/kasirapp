package repository

import (
	"database/sql"
	"log"
	"main/config"
	"main/model"
)

type ProductRepository interface {
	Create(payload model.Product) (model.Product, error)
}

type productRepositoy struct {
	db *sql.DB
}

func (p *productRepositoy) Create(payload model.Product) (model.Product, error) {
	var product model.Product

	err := p.db.QueryRow(config.InsertProduct, payload.Name, payload.Harga, payload.Stock).Scan(&product.ID, &product.Created_at)
	if err != nil {
		log.Println("productRepo.QueryRow:", err.Error())
		return model.Product{}, err
	}
	product.Name = payload.Name
	product.Harga = payload.Harga
	product.Stock = payload.Stock

	return product, nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepositoy{db: db}
}
