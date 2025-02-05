package db

import (
	"database/sql"
	"log"
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
