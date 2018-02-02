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
