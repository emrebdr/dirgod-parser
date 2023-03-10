package parser

import (
	"log"

	"github.com/emrebdr/dirgod-parser/src/constants"
	"github.com/emrebdr/dirgod-parser/src/models"
)

func (p *Parser) Literal() *models.Token {
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
	}

	log.Panicf("Unexpected token %s at line %d", p.nextLine.Value, p.CurrentLine+1)
	return nil
}

func (p *Parser) fromLiteral() *models.Token {
	token := p.Eat(constants.FROM)
	return &models.Token{
		Type:  constants.FROM,
		Value: token.Value,
	}
}

func (p *Parser) createFolderLiteral() *models.Token {
	token := p.Eat(constants.CREATEFOLDER)
	return &models.Token{
		Type:  constants.CREATEFOLDER,
		Value: token.Value,
	}
}

func (p *Parser) createFileLiteral() *models.Token {
	token := p.Eat(constants.CREATEFILE)
	return &models.Token{
		Type:  constants.CREATEFILE,
		Value: token.Value,
	}
}

func (p *Parser) workDirLiteral() *models.Token {
	token := p.Eat(constants.WORKDIR)
	return &models.Token{
		Type:  constants.WORKDIR,
		Value: token.Value,
	}
}

func (p *Parser) copyFolderLiteral() *models.Token {
	token := p.Eat(constants.COPYFOLDER)
	return &models.Token{
		Type:  constants.COPYFOLDER,
		Value: token.Value,
	}
}

func (p *Parser) copyFileLiteral() *models.Token {
	token := p.Eat(constants.COPYFILE)
	return &models.Token{
		Type:  constants.COPYFILE,
		Value: token.Value,
	}
}

func (p *Parser) moveFileLiteral() *models.Token {
	token := p.Eat(constants.MOVEFILE)
	return &models.Token{
		Type:  constants.MOVEFILE,
		Value: token.Value,
	}
}

func (p *Parser) moveFolderLiteral() *models.Token {
	token := p.Eat(constants.MOVEFOLDER)
	return &models.Token{
		Type:  constants.MOVEFOLDER,
		Value: token.Value,
	}
}

func (p *Parser) deleteFolderLiteral() *models.Token {
	token := p.Eat(constants.DELETEFOLDER)
	return &models.Token{
		Type:  constants.DELETEFOLDER,
		Value: token.Value,
	}
}

func (p *Parser) deleteFileLiteral() *models.Token {
	token := p.Eat(constants.DELETEFILE)
	return &models.Token{
		Type:  constants.DELETEFILE,
		Value: token.Value,
	}
}

func (p *Parser) chmodLiteral() *models.Token {
	token := p.Eat(constants.CHMOD)
	return &models.Token{
		Type:  constants.CHMOD,
		Value: token.Value,
	}
}

func (p *Parser) chownLiteral() *models.Token {
	token := p.Eat(constants.CHOWN)
	return &models.Token{
		Type:  constants.CHOWN,
		Value: token.Value,
	}
}
