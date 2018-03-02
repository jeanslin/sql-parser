# Golang SQL Parser

## Description
This lib need for parsing big sql file or just string row with comments and SQL request. You'll get a string slice with SQL requests without comments, tabs, new lines and excess spaces.

## Functions

### ParseFromFile
Input data: (filename string)
Output data: (requests []string, err error)

This function returns a string slice of SQL requests without comments ('#', '--', '/**/'), '\t', '\n' and excess spaces. The file is read by 1 mb to reduce loading of memory.

### ParseFromString
Input data: (input string)
Output data: (requests []string, err error)

This function receives just not formatted string row with SQL requests. It returns a string slice of SQL requests without comments ('#', '--', '/**/'), '\t', '\n' and excess spaces.# sql-parser

### Usage
```
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

/*Received queries*/


/*DROP TABLE IF EXISTS `sometesttable`;
__________________________________________________________________________
CREATE TABLE `sometesttable` ( `id` int(11) NOT NULL AUTO_INCREMENT, `Field_1` int(11) NOT NULL DEFAULT '0', `Field_211111112` varchar(11) NOT NULL DEFAULT '', `Field_3` int(11) NOT NULL DEFAULT '0', `Field_4` int(11) NOT NULL DEFAULT '0', `Field_5` int(11) NOT NULL DEFAULT '0', `Field_6` int(11) NOT NULL DEFAULT '0', PRIMARY KEY (`id`) ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
__________________________________________________________________________
LOCK TABLES `sometesttable` WRITE;
__________________________________________________________________________
INSERT INTO `sometesttable` (`id`, `Field_1`, `Field_211111112`, `Field_3`, `Field_4`, `Field_5`, `Field_6`) VALUES (1,0,'0',0,0,0,0), (2,0,'something',0,0,0,0), (3,0,'something',0,0,0,0), (4,0,'something',0,0,0,0), (5,0,'something',0,0,0,0);
__________________________________________________________________________
UNLOCK TABLES;*/
```