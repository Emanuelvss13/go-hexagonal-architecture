package application_test

import (
	"testing"

	application "github.com/emanuelvss13/go-hexagonal/app"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable_WhenPriceIsGreaterThanZero(t *testing.T) {
	product := application.Product{}
	product.Name = "Product"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)
}

func TestProduct_Enable_WhenPriceIsLessThanZero(t *testing.T) {
	product := application.Product{}
	product.Name = "Product"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Enable()

	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable_WhenPriceIsZero(t *testing.T) {
	product := application.Product{}
	product.Name = "Product"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)
}

func TestProduct_Disable_WhenPriceGreaterThanZero(t *testing.T) {
	product := application.Product{}
	product.Name = "Product"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Disable()

	require.Equal(t, "the price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = "c9e5c664-79f2-4dd2-b463-fc67de4389ae"
	product.Name = "Product"
	product.Status = application.DISABLED
	product.Price = 10

	bool, err := product.IsValid()

	require.True(t, bool)
	require.Nil(t, err)

	product.Status = "INVALID"

	bool, err = product.IsValid()

	require.False(t, bool)
	require.Equal(t, "product status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED

	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10

	_, err = product.IsValid()

	require.Equal(t, "product price must be greater than or equal zero", err.Error())

}
