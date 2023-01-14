package cli

import (
	"fmt"

	application "github.com/emanuelvss13/go-hexagonal/app"
)

func Run(service application.ProductServiceInterface, action string, id string, name string, price float64) (string, error) {
	result := ""

	switch action {
	case "create":
		product, err := service.Create(name, price)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	case "enable":
		product, err := service.Get(id)

		if err != nil {
			return result, err
		}

		res, err := service.Enable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s has been enabled", res.GetName())

	case "disable":
		product, err := service.Get(id)

		if err != nil {
			return result, err
		}

		res, err := service.Disable(product)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s has been disabled", res.GetName())

	default:
		product, err := service.Get(id)

		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	}

	return result, nil
}
