package models

import (
	"github.com/emrebdr/dirgod-parser/src/constants"
)

type Token struct {
	Type  constants.TokenType
	Value string
}
