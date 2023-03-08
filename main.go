package main

import (
	"encoding/csv"
	"fmt"
	"os"
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
	fmt.Println(records[0])
}
