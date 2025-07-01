package parser

import (
	"fmt"
	"regexp"

	"github.com/emrebdr/dirgod-parser/src/constants"
	"github.com/emrebdr/dirgod-parser/src/models"
)

type Tokenizer struct {
	Lines       []string
	CurrentLine uint16
}

func NewTokenizer(Lines []string) *Tokenizer {
	return &Tokenizer{
		Lines:       Lines,
		CurrentLine: 0,
	}
}

func (t *Tokenizer) HasNext() bool {
	return len(t.Lines) > int(t.CurrentLine)
}

func (t *Tokenizer) Next() (*models.Token, error) {
	if !t.HasNext() {
		return nil, nil
	}

	line := t.Lines[t.CurrentLine]
	for tokenType, regex := range constants.TOKENS {
		token := t.matchToken(regex, line)
		if token == "" {
			continue
		}

		if tokenType == constants.COMMENTLINE {
			return t.Next()
		}

		return &models.Token{
			Type:  tokenType,
			Value: token,
		}, nil
	}

	return nil, fmt.Errorf("error parsing line %d: %s", t.CurrentLine+1, line)
}

func (t *Tokenizer) matchToken(regex *regexp.Regexp, line string) string {
	token := regex.FindString(line)
	if token == "" {
		return ""
	}
	t.CurrentLine++
	return token
}
