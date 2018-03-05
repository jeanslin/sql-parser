package sqlparser

import (
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"
)

const (
	ErrorOpenFile = 1
	ErrorReadFile = 2
)

var (
	quote1, quote2, quote3                                       int
	isComment, isMultiComment, firstMinus, firstSlash, firstStar bool
	buf                                                          string
)

type Parser struct {
	sync.Mutex
}

type Error struct {
	Type    int
	Message string
}

func newParserError(code int, message string) Error {
	return Error{
		Type:    code,
		Message: message,
	}
}

func (e Error) Error() string {
	return strconv.Itoa(e.Type) + ": " + e.Message
}

/*
ParseFromFile - parse SQL-file and get requests type []string
*/
func (p *Parser) ParseFromFile(filename string) ([]string, error) {
	p.Lock()
	defer p.Unlock()
	var result, reqs []string
	// var buf string
	file, err := os.Open(filename)
	if err != nil {
		return nil, newParserError(ErrorOpenFile, err.Error())
	}
	defer file.Close()

	for {
		textByte := make([]byte, 1<<20)
		length, err := file.Read(textByte)
		if length == 0 {
			break
		}
		if err != nil {
			return result, newParserError(ErrorReadFile, err.Error())
		}
		text := string(textByte)
		reqs = queryBuilder(text, length)
		result = append(result, reqs...)

	}
	cleanGlobals()
	return choreRequests(result), nil
}

/*
ParseFromString - parse SQL-string and get requests type []string
*/
func (p *Parser) ParseFromString(requests string) ([]string, error) {
	p.Lock()
	defer p.Unlock()
	result := queryBuilder(requests, len(requests))
	cleanGlobals()
	return choreRequests(result), nil
}

func queryBuilder(text string, length int) []string {
	var requests []string
	var req string
	l := utf8.RuneCountInString(text)
	for i := 0; i < l; i++ {
		char := string([]rune(text)[i])
		// Check beginning the comment
		if strings.EqualFold(char, "-") {
			if quote1%2 == 0 && quote2%2 == 0 && quote3%2 == 0 {
				if firstMinus {
					isComment = true
					req = deleteLastSymbol(req)
					firstMinus = false
				} else {
					firstMinus = true
				}
			}
		} else {
			firstMinus = false
		}
		if strings.EqualFold(char, "*") {
			if quote1%2 == 0 && quote2%2 == 0 && quote3%2 == 0 {
				if firstSlash {
					isMultiComment = true
					req = deleteLastSymbol(req)
					firstSlash = false
				}
				firstStar = true
			}
		}
		if strings.EqualFold(char, "#") {
			if quote1%2 == 0 && quote2%2 == 0 && quote3%2 == 0 {
				isComment = true
			}
		}
		// Write char into req, if string is not commented
		if !isComment && !isMultiComment {
			// Find quotes
			if strings.EqualFold(char, "\"") {
				quote1++
			}
			if strings.EqualFold(char, "'") {
				quote2++
			}
			if strings.EqualFold(char, "`") {
				quote3++
			}
			req += char
		}
		// Turn off comment
		if strings.EqualFold(char, "/") {
			if quote1%2 == 0 && quote2%2 == 0 && quote3%2 == 0 {
				if firstStar {
					isMultiComment = false
					firstStar = false
				}
				firstSlash = true
			}
		}

		if isComment && char == "\n" {
			isComment = false
		}
		// Split request, if current char == ";" and no open quotes, no comments
		if quote1%2 == 0 && quote2%2 == 0 && quote3%2 == 0 && !isComment && !isMultiComment {
			if strings.EqualFold(char, ";") {
				requests = append(requests, buf+req)
				buf = ""
				req = ""
			}
			quote1, quote2, quote3 = 0, 0, 0
		}
	}
	buf += req
	return requests
}

func choreRequests(input []string) (requests []string) {
	for _, item := range input {
		item = strings.Replace(item, "\n\n", "\n", -1)
		item = strings.Replace(item, "\r\n", "\n", -1)
		item = strings.Replace(item, "\n", " ", -1)
		item = strings.Replace(item, "\t", " ", -1)

		for strings.Index(item, "  ") != -1 {
			item = strings.Replace(item, "  ", " ", -1)
		}
		for strings.Index(item, ";") == 0 || strings.Index(item, " ") == 0 {
			if strings.Index(item, ";") == 0 {
				item = strings.Replace(item, ";", "", 1)
			}
			if strings.Index(item, " ") == 0 {
				item = strings.Replace(item, " ", "", 1)
			}
		}
		if len(item) > 5 {
			requests = append(requests, item)
		}
	}
	return
}

func deleteLastSymbol(str string) string {
	if len(str) <= 1 {
		return ""
	}
	s := str[:len(str)-1]
	return s
}

func cleanGlobals() {
	quote1, quote2, quote3 = 0, 0, 0
	isComment, isMultiComment, firstMinus, firstSlash, firstStar = false, false, false, false, false
	buf = ""
}
