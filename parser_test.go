package sqlparser

import (
	"fmt"
	"testing"
)

/*
TestNewParserError - unit test for constructor parser error
*/
func TestNewParserError(t *testing.T) {
	var err error
	err = NewParserError(1, "error read file")
	if err != nil {
		if e, ok := err.(Error); !ok {
			t.Error("Error: error type is not valid!", e)
		}
	}
}

/*
TestParseFromFile - unit test for parsing sql requests from file
*/
func TestError(t *testing.T) {
	var err error
	err = NewParserError(1, "error read file")
	if err != nil {
		if e, ok := err.(Error); ok {
			if e.Error() == "" {
				t.Error("Error: function error() returns empty string!")
			}
		} else {
			t.Error("Error: error type is not valid!", e)
		}
	}
}

/*
TestParseFromFile - unit test for parsing sql requests from file
*/
func TestParseFromFile(t *testing.T) {
	result, err := ParseFromFile("asdasdas")
	if e, ok := err.(Error); ok {
		if result != nil || e.Type != 1 {
			t.Error("Error: Invalid result with incorrect file!")
		}
	}
	result, err = ParseFromFile("asdasdas")
	if e, ok := err.(Error); ok {
		if result != nil || e.Type != 1 {
			t.Error("Error: Invalid result with incorrect file!")
		}
	}
	result, err = ParseFromFile("test/test.sql")
	if err != nil {
		if e, ok := err.(Error); ok {
			if e.Type != ErrorOpenFile && e.Type != ErrorReadFile {
				t.Error("Error: ", e.Message)
			}
		} else {
			t.Error("Error: error type is not valid!")
		}
	}
	if len(result) == 0 {
		t.Error("Error: Result is empty!")
	}
	if result[0] != "INSERT INTO market_data.instruments (id, symbol, name, lot_size, lot_step, lot_max, lot_min, base_currency, price_currency, stop_level, state_tick, state_minute, symbol_type, swap_long, swap_short, fee, digits, point) VALUES (736, 'MCO;.N', 'Moody', 1, 1, 10000, 1, 'USD', '', 0.03, 'verifying', 'verifying', 'Equities USA', -1.13, -0.49, 0.1, 2, 0.01);" {
		t.Log(result[0])
		t.Error("Wrong result[0]!")
	}
	if result[1] != "INSERT INTO market_data.instruments (id, symbol, name, lot_size, lot_step, lot_max, lot_min, base_currency, price_currency, stop_level, state_tick, state_minute, symbol_type, swap_long, swap_short, fee, digits, point) VALUES (740, 'WDIG.DE', 'Wire Card', 1, 1, 10000, 1, 'EUR', '', 0.003, 'verifying', 'verifying', 'Equities DE', -3.08, -4.22, 0.1, 3, 0.001);" {
		t.Log(result[1])
		t.Error("Wrong result[1]!")
	}
}

/*
TestParseFromString - unit test for parsing sql requests from string
*/
func TestParseFromString(t *testing.T) {
	var input string
	input = "-- ;;;;;;\n#Comment is her;e\n/*And; ne   xt; multi\nline com\nme;nt\n\n\n*/;;;;;;;;INSERT INTO mar/*COMMENT  IS  HERE TOO*/ket_data.instruments \n\n\t (id, symbol, name, lot_size, lot_step, lot_max, lot_min, base_currency, price_currency, stop_level, state_tick, state_minute, symbol_type, swap_long, swap_short, fee, digits, point) VALUES (736, 'MCO.N;', \"Moody\", 1, 1, 10000, 1, 'USD', '', 0.03, 'verifying', 'verifying', `Equities USA`, -1.13, -0.49, 0.1, 2, 0.01);INSERT INTO market_data.instruments (id, symbol, name, lot_size, lot_step, lot_max, lot_min, base_currency, price_currency, stop_level, state_tick, state_minute, symbol_type, swap_long, swap_short, fee, digits, point) VALUES (740, 'WDIG.DE', 'Wire Card', 1, 1, 10000, 1, 'EUR', '', 0.003, 'verifying', 'verifying', 'Equities DE', -3.08, -4.22, 0.1, 3, 0.001);"
	result, err := ParseFromString(input)
	if err != nil {
		if e, ok := err.(Error); ok {
			if e.Type != ErrorOpenFile && e.Type != ErrorReadFile {
				t.Error("Error: ", e.Message)
			}
		} else {
			t.Error("Error: error type is not valid!")
		}
	}
	if len(result) == 0 {
		t.Error("Error: Result is empty!")
	}
	if result[0] != "INSERT INTO market_data.instruments (id, symbol, name, lot_size, lot_step, lot_max, lot_min, base_currency, price_currency, stop_level, state_tick, state_minute, symbol_type, swap_long, swap_short, fee, digits, point) VALUES (736, 'MCO.N;', \"Moody\", 1, 1, 10000, 1, 'USD', '', 0.03, 'verifying', 'verifying', `Equities USA`, -1.13, -0.49, 0.1, 2, 0.01);" {
		t.Log(result[0])
		t.Error("Wrong result[0]!")
	}
	if result[1] != "INSERT INTO market_data.instruments (id, symbol, name, lot_size, lot_step, lot_max, lot_min, base_currency, price_currency, stop_level, state_tick, state_minute, symbol_type, swap_long, swap_short, fee, digits, point) VALUES (740, 'WDIG.DE', 'Wire Card', 1, 1, 10000, 1, 'EUR', '', 0.003, 'verifying', 'verifying', 'Equities DE', -3.08, -4.22, 0.1, 3, 0.001);" {
		t.Log(result[1])
		t.Error("Wrong result[1]!")
	}
}

/*
TestQueryBuilder - unit test for parsing sql requests
*/
func TestQueryBuilder(t *testing.T) {
	var input string
	input = "-- ;;;;;;\n#Comment is her;e\n/*And; ne   xt; multi\nline com\nme;nt\n\n\n*/INSERT INTO mar/*COMMENT IS HERE TOO*/ket_data.instruments (id, symbol, name, lot_size, lot_step, lot_max, lot_min, base_currency, price_currency, stop_level, state_tick, state_minute, symbol_type, swap_long, swap_short, fee, digits, point) VALUES (736, 'MCO.N;', 'Moody', 1, 1, 10000, 1, 'USD', '', 0.03, 'verifying', 'verifying', 'Equities USA', -1.13, -0.49, 0.1, 2, 0.01);INSERT INTO market_data.instruments (id, symbol, name, lot_size, lot_step, lot_max, lot_min, base_currency, price_currency, stop_level, state_tick, state_minute, symbol_type, swap_long, swap_short, fee, digits, point) VALUES (740, 'WDIG.DE', 'Wire Card', 1, 1, 10000, 1, 'EUR', '', 0.003, 'verifying', 'verifying', 'Equities DE', -3.08, -4.22, 0.1, 3, 0.001);"
	result := queryBuilder(input, len(input))
	if len(result) == 0 {
		t.Error("Error: Result is empty!")
	}
	if result[0] != "INSERT INTO market_data.instruments (id, symbol, name, lot_size, lot_step, lot_max, lot_min, base_currency, price_currency, stop_level, state_tick, state_minute, symbol_type, swap_long, swap_short, fee, digits, point) VALUES (736, 'MCO.N;', 'Moody', 1, 1, 10000, 1, 'USD', '', 0.03, 'verifying', 'verifying', 'Equities USA', -1.13, -0.49, 0.1, 2, 0.01);" {
		fmt.Println(result[0])
		t.Error("Wrong result[0]!")
	}
	if result[1] != "INSERT INTO market_data.instruments (id, symbol, name, lot_size, lot_step, lot_max, lot_min, base_currency, price_currency, stop_level, state_tick, state_minute, symbol_type, swap_long, swap_short, fee, digits, point) VALUES (740, 'WDIG.DE', 'Wire Card', 1, 1, 10000, 1, 'EUR', '', 0.003, 'verifying', 'verifying', 'Equities DE', -3.08, -4.22, 0.1, 3, 0.001);" {
		fmt.Println(result[1])
		t.Error("Wrong result[1]!")
	}
	if buf != "" {
		t.Error("Wrong output data! buf is not empty!")
	}
}

/*
TestChoreRequests - unit test for deleting spaces, tabs and new lines
*/
func TestChoreRequests(t *testing.T) {
	var input []string
	input = append(input, "          		select  * from market_data.minutes  	where symbol = 'NEXT.PA'\n\nLimit	10;")
	input = append(input, "\n		select  *	\t	 from market_data.instruments  	where 	symbol = 'S.N'\n\nLimit  	\n	10;")
	result := choreRequests(input)
	if result[0] != "select * from market_data.minutes where symbol = 'NEXT.PA' Limit 10;" {
		t.Error("Wrong result[0]!")
	}
	if result[1] != "select * from market_data.instruments where symbol = 'S.N' Limit 10;" {
		t.Error("Wrong result[1]!")
	}
}
