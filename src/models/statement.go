package models

import "github.com/emrebdr/dirgod-parser/src/constants"

type StatementList struct {
	Statement []*Statement
}

type Statement struct {
	Type constants.TokenType
	Body interface{}
}
