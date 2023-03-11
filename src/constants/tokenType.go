package constants

type TokenType string

const (
	// Token types
	FROM                TokenType = "FROM"
	EXPRESSIONSTATEMENT TokenType = "EXPRESSIONSTATEMENT"
	BLOCKSTATEMENT      TokenType = "BLOCKSTATEMENT"
	PROGRAM             TokenType = "PROGRAM"
	WORKDIR             TokenType = "WORKDIR"
	COMMENTLINE         TokenType = "COMMENTLINE"
	CREATEFOLDER        TokenType = "CREATEFOLDER"
	CREATEFILE          TokenType = "CREATEFILE"
	DELETEFOLDER        TokenType = "DELETEFOLDER"
	DELETEFILE          TokenType = "DELETEFILE"
	COPYFOLDER          TokenType = "COPYFOLDER"
	COPYFILE            TokenType = "COPYFILE"
	MOVEFOLDER          TokenType = "MOVEFOLDER"
	MOVEFILE            TokenType = "MOVEFILE"
	CHMOD               TokenType = "CHMOD"
	CHOWN               TokenType = "CHOWN"
	EOF                 TokenType = "EOF"
)
