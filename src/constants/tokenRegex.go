package constants

import "regexp"

var FROMREGEX = regexp.MustCompile(`^FROM(?:\s+\S+(?:\/\S+)*){1}$`)
var WORKDIRREGEX = regexp.MustCompile(`^WORKDIR(?:\s+\S+(?:\/\S+)*){1}$`)
var COMMENTLINEREGEX = regexp.MustCompile(`^\/\/.*`)
var CREATEFOLDERREGEX = regexp.MustCompile(`^CREATEFOLDER(?:\s+\S+(?:\/\S+)*)+$`)
var CREATEFILEREGEX = regexp.MustCompile(`^CREATEFILE(?:\s+\S+(?:\/\S+)*)+$`)
var DELETEFOLDERREGEX = regexp.MustCompile(`^DELETEFOLDER(?:\s+\S+(?:\/\S+)*)+$`)
var DELETEFILEREGEX = regexp.MustCompile(`^DELETEFILE(?:\s+\S+(?:\/\S+)*)+$`)
var COPYFOLDERREGEX = regexp.MustCompile(`^COPYFOLDER(?:\s+\S+(?:\/\S+)*){2}$`)
var COPYFILEREGEX = regexp.MustCompile(`^COPYFILE(?:\s+\S+(?:\/\S+)*){2}$`)
var MOVEFOLDERREGEX = regexp.MustCompile(`^MOVEFOLDER(?:\s+\S+(?:\/\S+)*){2}$`)
var MOVEFILEREGEX = regexp.MustCompile(`^MOVEFILE(?:\s+\S+(?:\/\S+)*){2}$`)
var CHMODREGEX = regexp.MustCompile(`^CHMOD(?:\s+\S+(?:\/\S+)*){2,}$`)
var CHOWNREGEX = regexp.MustCompile(`^CHOWN(?:\s+\S+(?:\/\S+)*){2,}$`)
