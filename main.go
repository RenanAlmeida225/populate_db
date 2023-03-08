package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ReadCsv("test.csv")
}

func ReadCsv(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err.Error())
	}
	CreateTable(path, records[0])
}

func Db() *sql.DB {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func CreateTable(tableName string, columns []string) {
	var column string
	for i := 0; i < len(columns); i++ {
		if i == 0 {
			column += fmt.Sprintf("%s TEXT PRIMARY KEY NOT NULL, ", columns[i])
		} else if i == len(columns)-1 {
			column += fmt.Sprintf("%s TEXT", columns[i])
		} else {
			column += fmt.Sprintf("%s TEXT, ", columns[i])
		}
	}
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", strings.Split(tableName, ".csv")[0], column)
	db := Db()
	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	stmt.Exec()
	defer db.Close()
	fmt.Println(query)
}

// func InsertDb() {

// }
