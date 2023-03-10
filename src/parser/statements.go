package parser

import (
	"github.com/emrebdr/dirgod-parser/src/constants"
	"github.com/emrebdr/dirgod-parser/src/models"
)

func (p *Parser) StatementList(stopDelimiter constants.TokenType) *models.StatementList {
	statementList := &models.StatementList{
		Statement: []*models.Statement{},
	}
	if stopDelimiter == constants.EOF {
		for p.nextLine != nil {
			statementList.Statement = append(statementList.Statement, p.Statement())
		}
	} else {
		for p.nextLine != nil && p.nextLine.Type != stopDelimiter {
			statementList.Statement = append(statementList.Statement, p.Statement())
		}
	}

	return statementList
}

func (p *Parser) Statement() *models.Statement {
	switch p.nextLine.Type {
	case constants.FROM:
		return p.BlockStatement()
	default:
		return p.ExpressionStatement()
	}
}

func (p *Parser) ExpressionStatement() *models.Statement {
	expression := p.Expression()
	return &models.Statement{
		Type: constants.EXPRESSIONSTATEMENT,
		Body: *expression,
	}
}

func (p *Parser) Expression() *models.Token {
	return p.Literal()
}

func (p *Parser) BlockStatement() *models.Statement {
	p.Eat(constants.FROM)

	var body *models.StatementList
	if p.nextLine.Type != constants.FROM {
		body = p.StatementList(constants.FROM)
	} else {
		body = &models.StatementList{}
	}

	return &models.Statement{
		Type: constants.BLOCKSTATEMENT,
		Body: body,
	}
}
