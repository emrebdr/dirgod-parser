package parser

import (
	"fmt"
	"github.com/emrebdr/dirgod-parser/src/constants"
	"github.com/emrebdr/dirgod-parser/src/models"
)

func (p *Parser) Literal() (*models.Token, error) {
	switch p.nextLine.Type {
	case constants.CREATEFOLDER:
		return p.createFolderLiteral()
	case constants.CREATEFILE:
		return p.createFileLiteral()
	case constants.WORKDIR:
		return p.workDirLiteral()
	case constants.COPYFILE:
		return p.copyFileLiteral()
	case constants.COPYFOLDER:
		return p.copyFolderLiteral()
	case constants.MOVEFILE:
		return p.moveFileLiteral()
	case constants.MOVEFOLDER:
		return p.moveFolderLiteral()
	case constants.DELETEFILE:
		return p.deleteFileLiteral()
	case constants.DELETEFOLDER:
		return p.deleteFolderLiteral()
	case constants.CHMOD:
		return p.chmodLiteral()
	case constants.CHOWN:
		return p.chownLiteral()
	case constants.FROM:
		return p.fromLiteral()
	case constants.WRITE:
		return p.writeLiteral()
	}

	return nil, fmt.Errorf("unexpected token %s at line %d", p.nextLine.Value, p.CurrentLine+1)
}

func (p *Parser) fromLiteral() (*models.Token, error) {
	token, err := p.Eat(constants.FROM)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		Type:  constants.FROM,
		Value: token.Value,
	}, nil
}

func (p *Parser) createFolderLiteral() (*models.Token, error) {
	token, err := p.Eat(constants.CREATEFOLDER)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		Type:  constants.CREATEFOLDER,
		Value: token.Value,
	}, nil
}

func (p *Parser) createFileLiteral() (*models.Token, error) {
	token, err := p.Eat(constants.CREATEFILE)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		Type:  constants.CREATEFILE,
		Value: token.Value,
	}, nil
}

func (p *Parser) workDirLiteral() (*models.Token, error) {
	token, err := p.Eat(constants.WORKDIR)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		Type:  constants.WORKDIR,
		Value: token.Value,
	}, nil
}

func (p *Parser) copyFolderLiteral() (*models.Token, error) {
	token, err := p.Eat(constants.COPYFOLDER)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		Type:  constants.COPYFOLDER,
		Value: token.Value,
	}, nil
}

func (p *Parser) copyFileLiteral() (*models.Token, error) {
	token, err := p.Eat(constants.COPYFILE)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		Type:  constants.COPYFILE,
		Value: token.Value,
	}, nil
}

func (p *Parser) moveFileLiteral() (*models.Token, error) {
	token, err := p.Eat(constants.MOVEFILE)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		Type:  constants.MOVEFILE,
		Value: token.Value,
	}, nil
}

func (p *Parser) moveFolderLiteral() (*models.Token, error) {
	token, err := p.Eat(constants.MOVEFOLDER)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		Type:  constants.MOVEFOLDER,
		Value: token.Value,
	}, nil
}

func (p *Parser) deleteFolderLiteral() (*models.Token, error) {
	token, err := p.Eat(constants.DELETEFOLDER)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		Type:  constants.DELETEFOLDER,
		Value: token.Value,
	}, nil
}

func (p *Parser) deleteFileLiteral() (*models.Token, error) {
	token, err := p.Eat(constants.DELETEFILE)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		Type:  constants.DELETEFILE,
		Value: token.Value,
	}, nil
}

func (p *Parser) chmodLiteral() (*models.Token, error) {
	token, err := p.Eat(constants.CHMOD)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		Type:  constants.CHMOD,
		Value: token.Value,
	}, nil
}

func (p *Parser) chownLiteral() (*models.Token, error) {
	token, err := p.Eat(constants.CHOWN)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		Type:  constants.CHOWN,
		Value: token.Value,
	}, nil
}

func (p *Parser) writeLiteral() (*models.Token, error) {
	token, err := p.Eat(constants.WRITE)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		Type:  constants.WRITE,
		Value: token.Value,
	}, nil
}
