package main

import (
	"fmt"
	"log"

	sqlparser "github.com/PGV65/sql-parser"
)

func main() {
	var parser sqlparser.Parser
	// requests, err := parser.ParseFromString("/*Open\ntable\ntest*/insert into test.test (`id`, `count`, `comment`) values (14,10,'iopoip');\ninsert into test.test (`id`, `count`, `comment`) values (15,10,'1213ewqdsa')\n;-- Open table test again\nselect * from test.test;")
	requests, err := parser.ParseFromFile("../test/test.sql")
	if err != nil {
		log.Fatal(err)
	}
	for i := range requests {
		fmt.Println(requests[i])
		fmt.Println("__________________________________________________________________________")
	}
}
