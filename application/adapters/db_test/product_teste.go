package db

import (
	"database/sql"
	"log"
	"testing"

	"github.com/jadersoncardoso/codeedu/go-hexagonal/application/adapters/db"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
			"id" string
			"name" string
			"price" float
			"status" string
	)`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abc", "Product Teste", 0, "disabled)`
	smtp, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	smtp.Exec()

}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, "0.0", product.GetPrice())
	require.Equal(t, "Product Test", product.GetStatus())

}
