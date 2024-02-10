package repository

import (
	"database/sql"
	"log"
	"main/config"
	"main/model"
)

type ProductRepository interface {
	Create(payload model.Product) (model.Product, error)
	LIst() ([]model.Product, error)
}

type productRepositoy struct {
	db *sql.DB
}

// LIst implements ProductRepository.
func (r *productRepositoy) LIst() ([]model.Product, error) {
	var products []model.Product

	rows, err := r.db.Query(config.SelectProduct)
	if err != nil {
		log.Println("productRepository.Query:", err.Error())
		return []model.Product{}, err
	}

	for rows.Next() {
		var product model.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Harga,
			&product.Stock,
			&product.Created_at,
			&product.Updated_at,
		)

		if err != nil {
			log.Println("productRepository.Rows.Next:", err.Error())
			return []model.Product{}, err
		}

		products = append(products, product)
	}

	return products, nil
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
