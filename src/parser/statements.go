package parser

import (
	"github.com/emrebdr/dirgod-parser/src/constants"
	"github.com/emrebdr/dirgod-parser/src/models"
)

func (p *Parser) StatementList(stopDelimiter constants.TokenType) (*models.StatementList, error) {
	statementList := &models.StatementList{
		Statement: []models.Statement{},
	}
	if stopDelimiter == constants.EOF {
		for p.nextLine != nil {
			statement, err := p.Statement()
			if err != nil {
				return nil, err
			}
			statementList.Statement = append(statementList.Statement, statement)
		}
	} else {
		for p.nextLine != nil && p.nextLine.Type != stopDelimiter {
			statement, err := p.Statement()
			if err != nil {
				return nil, err
			}
			statementList.Statement = append(statementList.Statement, statement)
		}
	}

	return statementList, nil
}

func (p *Parser) Statement() (models.Statement, error) {
	switch p.nextLine.Type {
	case constants.FROM:
		return p.BlockStatement()
	default:
		return p.ExpressionStatement()
	}
}

func (p *Parser) ExpressionStatement() (models.Statement, error) {
	expression, err := p.Expression()
	if err != nil {
		return models.Statement{}, err
	}

	return models.Statement{
		Type: constants.EXPRESSIONSTATEMENT,
		Body: *expression,
	}, nil
}

func (p *Parser) Expression() (*models.Token, error) {
	return p.Literal()
}

func (p *Parser) BlockStatement() (models.Statement, error) {
	_, err := p.Eat(constants.FROM)
	if err != nil {
		return models.Statement{}, err
	}

	var body *models.StatementList
	if p.nextLine.Type != constants.FROM {
		body, err = p.StatementList(constants.FROM)
		if err != nil {
			return models.Statement{}, nil
		}
	} else {
		body = &models.StatementList{}
	}

	return models.Statement{
		Type: constants.BLOCKSTATEMENT,
		Body: body,
	}, nil
}
