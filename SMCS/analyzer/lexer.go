package analyzer

import (
	"github.com/csos95/CMPS-4113-Spring-2017/SMCS/analyzer/c"
	"log"
)

type Token struct {
	Type  string
	Value string
}
type NextToken func() (string, string)

func Tokenize(language, source string) []Token {
	var next NextToken
	var tokens []Token
	switch language {
	case "c":
		c.Parse(source)
		next = c.NextToken
	}

	ntoken, vtoken := next()

	for ntoken != "NULL" {
		tokens = append(tokens, Token{Type: ntoken, Value: vtoken})
		log.Println(ntoken, vtoken)
		ntoken, vtoken = next()
	}

	c.Close()

	return tokens
}
