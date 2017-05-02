package analyzer

import (
	"github.com/csos95/CMPS-4113-Spring-2017/SMCS/analyzer/CPP"
	"github.com/csos95/CMPS-4113-Spring-2017/SMCS/analyzer/Java"
	"github.com/csos95/CMPS-4113-Spring-2017/SMCS/analyzer/c"
	//"log"
)

type Token struct {
	Type  string
	Value string
}
type NextToken func() (string, string)
type Close func()

func Tokenize(language, source string) []Token {
	var next NextToken
	var tokens []Token
	var closeCall Close
	switch language {
	case "c":
		c.Parse(source)
		next = c.NextToken
		closeCall = c.Close
	case "cpp":
		cpp.Parse(source)
		next = cpp.NextToken
		closeCall = cpp.Close
	case "java":
		java.Parse(source)
		next = java.NextToken
		closeCall = java.Close
	}

	ntoken, vtoken := next()

	for ntoken != "NULL" {
		tokens = append(tokens, Token{Type: ntoken, Value: vtoken})
		//log.Println(ntoken, vtoken)
		ntoken, vtoken = next()
	}

	closeCall()

	return tokens
}
