package parser

import (
	"bufio"
	"log"
	"os"

	"github.com/emrebdr/dirgod-parser/src/constants"
	"github.com/emrebdr/dirgod-parser/src/models"
)

type Parser struct {
	Config      *os.File
	Tokenizer   *Tokenizer
	Lines       []string
	CurrentLine uint16
	nextLine    *models.Token
}

// NewParser creates a new Parser object
func NewParser(configPath string) *Parser {
	//! remove the configPath from the parameters. Because configPath dynamically find the directory of the config file
	config, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Error opening config file: %s", err)
	}
	defer config.Close()
	lines := ReadLines(config)

	return &Parser{
		Config:      config,
		Tokenizer:   NewTokenizer(lines),
		Lines:       lines,
		CurrentLine: 0,
	}
}

func ReadLines(config *os.File) []string {
	fileScanner := bufio.NewScanner(config)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		newLine := fileScanner.Text()
		if len(newLine) == 0 {
			continue
		}

		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func (p *Parser) Program() *models.Program {
	return &models.Program{
		Type: constants.PROGRAM,
		Body: p.StatementList(constants.EOF),
	}
}

func (p *Parser) Parse() *models.Program {
	p.nextLine = p.Tokenizer.Next()
	return p.Program()
}

func (p *Parser) Eat(tokenType constants.TokenType) *models.Token {
	token := p.nextLine

	if token == nil {
		log.Panicf("Unexpected EOF at line %d", p.CurrentLine+1)
	}

	if token.Type != tokenType {
		log.Panicf("Unexpected token %s at line %d", token.Value, p.CurrentLine+1)
	}

	p.nextLine = p.Tokenizer.Next()

	return token
}
