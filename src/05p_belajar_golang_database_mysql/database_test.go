package belajar_golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3308)/belajar_golang_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// gunakan DB
}
