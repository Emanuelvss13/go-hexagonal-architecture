package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/emanuelvss13/go-hexagonal/adapters/db"
	application "github.com/emanuelvss13/go-hexagonal/app"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);`

	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values ("test", "Test Product", 0, "disabled")`

	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProduct_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDB(Db)

	result, err := productDb.Get("test")

	require.Nil(t, err)
	require.Equal(t, "Test Product", result.GetName())
	require.Equal(t, 0.0, result.GetPrice())
	require.Equal(t, "disabled", result.GetStatus())
}

func TestProduct_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	product := application.NewProduct()

	productDb := db.NewProductDB(Db)

	createdProduct, err := productDb.Save(product)

	require.Nil(t, err)

	require.Equal(t, product.Name, createdProduct.GetName())
	require.Equal(t, product.Price, createdProduct.GetPrice())
	require.Equal(t, product.Status, createdProduct.GetStatus())

	product.Enable()

	updatedProduct, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Status, updatedProduct.GetStatus())
}
