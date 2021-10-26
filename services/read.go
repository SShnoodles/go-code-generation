package services

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-code-generation/models"
	"log"
)

func Read(t models.DbType) []models.Schema {
	var schemas []models.Schema
	switch t {
	case models.Postgres:
		postgres(schemas)
	case models.Mysql:
		mysql(schemas)
	case models.Oracle:
		oracle(schemas)
	}
	return schemas
}

func postgres(schemas []models.Schema) {
	connStr := "postgres://postgres:123456@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT nspname FROM pg_catalog.pg_namespace")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var schema models.Schema
		rows.Scan(&schema.Name)
		fmt.Println("schema >>>>>>", schema.Name)

		err := db.QueryRow("SELECT COALESCE(obj_description($1::regnamespace),'')", schema.Name).Scan(&schema.Comment)
		if err != nil{
			log.Fatal(err)
		}

		rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema=$1 AND table_type='BASE TABLE'", schema.Name)
		if err != nil {
			log.Fatal(err)
		}
		var tables []models.Table
		for rows.Next() {
			var table models.Table
			rows.Scan(&table.Name)
			fmt.Println("table >", table.Name)

			err := db.QueryRow("SELECT COALESCE(obj_description($1::regclass),'')", schema.Name + "." + table.Name).Scan(&table.Comment)
			if err != nil{
				log.Fatal(err)
			}
			tables = append(tables, table)
		}

		schema.Tables = tables
		schemas = append(schemas, schema)
	}
}

func mysql(schema []models.Schema) {

}

func oracle(schema []models.Schema) {
}
