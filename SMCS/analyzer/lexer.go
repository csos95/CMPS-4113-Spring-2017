package analyzer

import (
	"github.com/csos95/CMPS-4113-Spring-2017/SMCS/analyzer/CPP"
	"github.com/csos95/CMPS-4113-Spring-2017/SMCS/analyzer/Java"
	"github.com/csos95/CMPS-4113-Spring-2017/SMCS/analyzer/c"
	//"log"
)

//Token holds the type and value of a token
type Token struct {
	Type  string
	Value string
}

//NextToken is a function type that will return the next token type and value
type NextToken func() (string, string)

//Close is a function type that will close the lex parser and free up its memory
type Close func()

//Tokenize takes a string of source and the name of the language the source code is written in
func Tokenize(language, source string) []Token {
	var next NextToken
	var tokens []Token
	var closeCall Close

	//choose which parsers functions to use based on the language
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

	//go through and get all of the tokens
	for ntoken != "NULL" {
		tokens = append(tokens, Token{Type: ntoken, Value: vtoken})
		//log.Println(ntoken, vtoken)
		ntoken, vtoken = next()
	}

	//free up the parser resources
	closeCall()

	return tokens
}
