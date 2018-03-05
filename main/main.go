package main

import (
	"fmt"
	"log"

	sqlparser "github.com/PGV65/sql-parser"
)

func main() {
	var parser sqlparser.Parser
	requests, err := parser.ParseFromString(`INSERT INTO instruments (name, lot_size, id) VALUES ("SPA35#", 0, 111);`)
	// requests, err := parser.ParseFromFile("../test/test.sql")
	if err != nil {
		log.Fatal(err)
	}
	for i := range requests {
		fmt.Println(requests[i])
		fmt.Println("__________________________________________________________________________")
	}
}
