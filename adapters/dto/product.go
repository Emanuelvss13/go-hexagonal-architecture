package dto

import application "github.com/emanuelvss13/go-hexagonal/app"

type ProductDTO struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProduct() *ProductDTO {
	return &ProductDTO{}
}

func (p *ProductDTO) Bind(product *application.Product) (*application.Product, error) {

	if p.ID != "" {
		product.ID = p.ID
	}

	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status

	_, err := product.IsValid()

	if err != nil {
		return &application.Product{}, err
	}

	return product, nil

}
