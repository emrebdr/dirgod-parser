package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/emrebdr/dirgod-parser/src/constants"
	"github.com/emrebdr/dirgod-parser/src/models"
)

type Parser struct {
	Tokenizer   *Tokenizer
	Lines       []string
	CurrentLine uint16
	nextLine    *models.Token
}

// NewParser creates a new Parser object
func NewParser(configPath string) (*Parser, error) {
	config, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %w", err)
	}
	defer config.Close()
	lines := ReadLines(config)

	return &Parser{
		Tokenizer:   NewTokenizer(lines),
		Lines:       lines,
		CurrentLine: 0,
	}, nil
}

// NewParserFromString creates a new Parser object from string content
func NewParserFromString(content string) (*Parser, error) {
	lines, err := ReadLinesFromString(content)
	if err != nil {
		return nil, err
	}

	return &Parser{
		Tokenizer:   NewTokenizer(lines),
		Lines:       lines,
		CurrentLine: 0,
	}, nil
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

func ReadLinesFromString(content string) ([]string, error) {
	scanner := bufio.NewScanner(strings.NewReader(content))

	var lines []string
	for scanner.Scan() {
		newLine := scanner.Text()
		if len(newLine) == 0 {
			continue
		}

		lines = append(lines, newLine)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading string: %w", err)
	}

	return lines, nil
}

func (p *Parser) Program() (*models.Program, error) {
	body, err := p.StatementList(constants.EOF)
	if err != nil {
		return nil, err
	}

	return &models.Program{
		Type: constants.PROGRAM,
		Body: body,
	}, nil
}

func (p *Parser) Parse() (*models.Program, error) {
	nextLine, err := p.Tokenizer.Next()
	if err != nil {
		return nil, err
	}

	p.nextLine = nextLine

	return p.Program()
}

func (p *Parser) IsValid() bool {
	_, err := p.Parse()
	if err != nil {
		return false
	}

	return true
}

func (p *Parser) Eat(tokenType constants.TokenType) (*models.Token, error) {
	token := p.nextLine

	if token == nil {
		return nil, fmt.Errorf("unexpected EOF at line %d", p.CurrentLine+1)
	}

	if token.Type != tokenType {
		return nil, fmt.Errorf("unexpected token %s at line %d", token.Value, p.CurrentLine+1)
	}

	nextLine, err := p.Tokenizer.Next()
	if err != nil {
		return nil, err
	}

	p.nextLine = nextLine

	return token, nil
}
