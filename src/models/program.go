package models

import (
	"encoding/json"

	"github.com/emrebdr/dirgod-parser/src/constants"
)

type Program struct {
	Type constants.TokenType
	Body *StatementList
}

func (p *Program) ToString() string {
	str, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(str)
}
