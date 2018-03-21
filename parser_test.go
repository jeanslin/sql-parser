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
	err = newParserError(1, "error read file")
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
	err = newParserError(1, "error read file")
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
	var parser Parser
	result, err := parser.ParseFromFile("asdasdas")
	if e, ok := err.(Error); ok {
		if result != nil || e.Type != 1 {
			t.Error("Error: Invalid result with incorrect file!")
		}
	}
	result, err = parser.ParseFromFile("asdasdas")
	if e, ok := err.(Error); ok {
		if result != nil || e.Type != 1 {
			t.Error("Error: Invalid result with incorrect file!")
		}
	}
	result, err = parser.ParseFromFile("test/test.sql")
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
	if result[0] != "DROP TABLE IF EXISTS `sometesttable`;" {
		t.Log(result[0])
		t.Error("Wrong result[0]!")
	}
	if result[1] != "CREATE TABLE `sometesttable` ( `id` int(11) NOT NULL AUTO_INCREMENT, `Field_1` int(11) NOT NULL DEFAULT '0', `Field_211111112` varchar(11) NOT NULL DEFAULT '', `Field_3` int(11) NOT NULL DEFAULT '0', `Field_4` int(11) NOT NULL DEFAULT '0', `Field_5` int(11) NOT NULL DEFAULT '0', `Field_6` int(11) NOT NULL DEFAULT '0', PRIMARY KEY (`id`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8;" {
		t.Log(result[1])
		t.Error("Wrong result[1]!")
	}
	if result[2] != "LOCK TABLES `sometesttable` WRITE;" {
		t.Log(result[2])
		t.Error("Wrong result[2]!")
	}
	if result[3] != "INSERT INTO `sometesttable` (`id`, `Field_1`, `Field_211111112`, `Field_3`, `Field_4`, `Field_5`, `Field_6`) VALUES (1,0,'0',0,0,0,0), (2,0,'something',0,0,0,0), (3,0,'something',0,0,0,0), (4,0,'something',0,0,0,0), (5,0,'something',0,0,0,0);" {
		t.Log(result[3])
		t.Error("Wrong result[3]!")
	}

	if result[4] != "UNLOCK TABLES;" {
		t.Log(result[4])
		t.Error("Wrong result[4]!")
	}

	result, err = parser.ParseFromFile("test/test1.sql")
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
	if result[0] != "CREATE TYPE session_status AS ENUM ('new', 'finished', 'active', 'declined');" {
		t.Log(result[0])
		t.Error("Wrong result[0]!")
	}
	if result[1] != "CREATE TABLE sessions ( id VARCHAR(255) NOT NULL CONSTRAINT sessions_pkey PRIMARY KEY, creatorid INTEGER NOT NULL, abonentid INTEGER NOT NULL, status SESSION_STATUS DEFAULT 'new' :: SESSION_STATUS NOT NULL, createdat TIMESTAMP DEFAULT timezone('utc' :: TEXT, now()) NOT NULL, updatedat TIMESTAMP DEFAULT timezone('UTC' :: TEXT, now()) NOT NULL );" {
		t.Log(result[1])
		t.Error("Wrong result[1]!")
	}
	if result[2] != "CREATE UNIQUE INDEX sessions_id_uindex ON sessions (id);" {
		t.Log(result[2])
		t.Error("Wrong result[2]!")
	}
	if result[3] != "CREATE OR REPLACE FUNCTION trigger_upd_time() RETURNS TRIGGER AS $$ BEGIN NEW.updatedat = (NOW() AT TIME ZONE 'UTC'); RETURN NEW; END; $$ LANGUAGE plpgsql;" {
		t.Log(result[3])
		t.Error("Wrong result[3]!")
	}

	if result[4] != "CREATE TRIGGER set_upd_time BEFORE UPDATE ON sessions FOR EACH ROW EXECUTE PROCEDURE trigger_upd_time();" {
		t.Log(result[4])
		t.Error("Wrong result[4]!")
	}
}

/*
TestParseFromString - unit test for parsing sql requests from string
*/
func TestParseFromString(t *testing.T) {
	var parser Parser
	var input string
	input = "-- ;;;;;;\n#Comment is her;e\n/*And; ne   xt; multi\nline com\nme;nt\n\n\n*/;;;;;;;;INSERT INTO mar/*COMMENT  IS  HERE TOO*/ket_data.instruments \n\n\t (id, symbol, name, lot_size, lot_step, lot_max, lot_min, base_currency, price_currency, stop_level, state_tick, state_minute, symbol_type, swap_long, swap_short, fee, digits, point) VALUES (736, 'MCO.N;', \"Moody\", 1, 1, 10000, 1, 'USD', '', 0.03, 'verifying', 'verifying', `Equities USA`, -1.13, -0.49, 0.1, 2, 0.01);INSERT INTO market_data.instruments (id, symbol, name, lot_size, lot_step, lot_max, lot_min, base_currency, price_currency, stop_level, state_tick, state_minute, symbol_type, swap_long, swap_short, fee, digits, point) VALUES (740, 'WDIG.DE', 'Wire Card', 1, 1, 10000, 1, 'EUR', '', 0.003, 'verifying', 'verifying', 'Equities DE', -3.08, -4.22, 0.1, 3, 0.001);"
	result, err := parser.ParseFromString(input)
	if err != nil {
		if e, ok := err.(Error); ok {
			if e.Type != ErrorOpenFile && e.Type != ErrorReadFile {
				t.Error("Error: ", e.Message)
			}
		} else {
			t.Error("Error: error type is not valid!")
		}
	}
	result2, err := parser.ParseFromString(`INSERT INTO instruments (name, lot_size, id) VALUES ("SPA35#", 0, 111);`)
	if err != nil {
		if e, ok := err.(Error); ok {
			if e.Type != ErrorOpenFile && e.Type != ErrorReadFile {
				t.Error("Error: ", e.Message)
			}
		} else {
			t.Error("Error: error type is not valid!")
		}
	}
	result3, err := parser.ParseFromString(`INSERT INTO instruments (name, lot_size, id) VALUES ("SPA35--", 0, 111);`)
	if err != nil {
		if e, ok := err.(Error); ok {
			if e.Type != ErrorOpenFile && e.Type != ErrorReadFile {
				t.Error("Error: ", e.Message)
			}
		} else {
			t.Error("Error: error type is not valid!")
		}
	}
	result4, err := parser.ParseFromString(`INSERT INTO instruments (name, lot_size, id) VALUES ("/*SPA3*/5", 0, 111);`)
	if err != nil {
		if e, ok := err.(Error); ok {
			if e.Type != ErrorOpenFile && e.Type != ErrorReadFile {
				t.Error("Error: ", e.Message)
			}
		} else {
			t.Error("Error: error type is not valid!")
		}
	}

	result5, err := parser.ParseFromString(`INSERT INTO instruments (name, lot_size, id) VALUES ("/*SPA35", 0, 111);`)
	if err != nil {
		if e, ok := err.(Error); ok {
			if e.Type != ErrorOpenFile && e.Type != ErrorReadFile {
				t.Error("Error: ", e.Message)
			}
		} else {
			t.Error("Error: error type is not valid!")
		}
	}

	result6, err := parser.ParseFromString(`INSERT INTO instruments (name, lot_size, id) VALUES (/*"SPA35"*/, 0, 111);`)
	if err != nil {
		if e, ok := err.(Error); ok {
			if e.Type != ErrorOpenFile && e.Type != ErrorReadFile {
				t.Error("Error: ", e.Message)
			}
		} else {
			t.Error("Error: error type is not valid!")
		}
	}

	result7, err := parser.ParseFromString(`INSERT INTO instruments (name, lot_size, id) VALUES ("SPA35АБВВВВ", 0, 111);`)
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

	if len(result2) == 0 {
		t.Error("Error: Result is empty!")
	}
	if result2[0] != `INSERT INTO instruments (name, lot_size, id) VALUES ("SPA35#", 0, 111);` {
		t.Log(result2[0])
		t.Error("Wrong result2[0]!")
	}

	if len(result3) == 0 {
		t.Error("Error: Result is empty!")
	}
	if result3[0] != `INSERT INTO instruments (name, lot_size, id) VALUES ("SPA35--", 0, 111);` {
		t.Log(result3[0])
		t.Error("Wrong result3[0]!")
	}

	if len(result4) == 0 {
		t.Error("Error: Result is empty!")
	}
	if result4[0] != `INSERT INTO instruments (name, lot_size, id) VALUES ("/*SPA3*/5", 0, 111);` {
		t.Log(result4[0])
		t.Error("Wrong result4[0]!")
	}

	if len(result5) == 0 {
		t.Error("Error: Result is empty!")
	}
	if result5[0] != `INSERT INTO instruments (name, lot_size, id) VALUES ("/*SPA35", 0, 111);` {
		t.Log(result5[0])
		t.Error("Wrong result5[0]!")
	}

	if len(result6) == 0 {
		t.Error("Error: Result is empty!")
	}
	if result6[0] != `INSERT INTO instruments (name, lot_size, id) VALUES (, 0, 111);` {
		t.Log(result6[0])
		t.Error("Wrong result6[0]!")
	}

	if len(result7) == 0 {
		t.Error("Error: Result is empty!")
	}
	if result7[0] != `INSERT INTO instruments (name, lot_size, id) VALUES ("SPA35АБВВВВ", 0, 111);` {
		t.Log(result7[0])
		t.Error("Wrong result7[0]!")
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
