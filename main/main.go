package main

import (
	"fmt"
	"log"
	sqlparser "sql-parser"
)

func main() {
	// requests, err := sqlparser.ParseFromString("-- ;;;;;;\n#Comment is her;e\n/*And; ne   xt; multi\nline com\nme;nt\n\n\n*/;;;;;;;;INSERT INTO mar/*COMMENT  IS  HERE TOO*/ket_data.instruments \n\n\t (id, symbol, name, lot_size, lot_step, lot_max, lot_min, base_currency, price_currency, stop_level, state_tick, state_minute, symbol_type, swap_long, swap_short, fee, digits, point) VALUES (736, 'MCO.N;', 'Moody', 1, 1, 10000, 1, 'USD', '', 0.03, 'verifying', 'verifying', 'Equities USA', -1.13, -0.49, 0.1, 2, 0.01);INSERT INTO market_data.instruments (id, symbol, name, lot_size, lot_step, lot_max, lot_min, base_currency, price_currency, stop_level, state_tick, state_minute, symbol_type, swap_long, swap_short, fee, digits, point) VALUES (740, 'WDIG.DE', 'Wire Card', 1, 1, 10000, 1, 'EUR', '', 0.003, 'verifying', 'verifying', 'Equities DE', -3.08, -4.22, 0.1, 3, 0.001);")
	requests, err := sqlparser.ParseFromFile("../test/test.sql")
	if err != nil {
		log.Fatal(err)
	}
	for i := range requests {
		fmt.Println(requests[i])
		fmt.Println("__________________________________________________________________________")
	}
}
