# dirgod-parser

A small **demo** parser for the **Dirgod** DSL — a custom language that describes file and folder operations (create, copy, move, delete, chmod, chown, etc.). Written in Go for fun.

## What is Dirgod?

Dirgod is a line-based DSL. Each line is a single command. Think of it like a tiny spec file for “what to do with files and folders.” The parser only **validates** and **parses** the file into an AST; it does not execute the operations.

## Requirements

- **Go 1.22+**

## Quick start

```bash
# Clone and enter the repo
cd dirgod-parser

# Run (validates example/Dirgod and prints true/false)
go run .
```

## Project structure

```
.
├── main.go              # Entry: parses example/Dirgod and prints IsValid()
├── go.mod
├── example/
│   └── Dirgod           # Sample Dirgod script
└── src/
    ├── constants/       # Token types and regex patterns
    ├── models/          # AST: Program, Statement, Token
    └── parser/          # Tokenizer + parser (line-by-line)
```

## Dirgod syntax (supported commands)

| Command       | Example |
|---------------|---------|
| `FROM`        | `FROM scratch` |
| `WORKDIR`     | `WORKDIR /var/www/html` |
| `CREATEFOLDER`| `CREATEFOLDER a/b/c` or `CREATEFOLDER a/b/c d/f/g` |
| `CREATEFILE`  | `CREATEFILE a/b/c/d.txt` or multiple paths |
| `DELETEFOLDER`| `DELETEFOLDER a/b/c` or multiple paths |
| `DELETEFILE`  | `DELETEFILE a/b/c.txt a/b/d.txt` |
| `COPYFOLDER`  | `COPYFOLDER a/b/c d/f/g` (source, dest) |
| `COPYFILE`    | `COPYFILE a/b/c.txt a/b/c` |
| `MOVEFOLDER`  | `MOVEFOLDER a/b/c d/f/g` |
| `MOVEFILE`    | `MOVEFILE a/b/c.txt a/b/c` |
| `CHMOD`       | `CHMOD 777 a/b/c.txt` or `CHMOD 777 a/b/c recursive` |
| `CHOWN`       | `CHOWN www-data:www-data a/b/c.txt` or `... recursive` |
| `WRITE`       | `WRITE <source> <dest> url` (e.g. URL to file path) |
| Comments      | `// comment` (full line) |

- Empty lines are skipped.
- Paths use `/` (e.g. `a/b/c`). The parser does not run any of these commands; it only checks that the script is valid.

## Usage in code

```go
package main

import (
    "fmt"
    "github.com/emrebdr/dirgod-parser/src/parser"
)

func main() {
    // From file
    p, err := parser.NewParser("./example/Dirgod")
    if err != nil {
        panic(err)
    }
    fmt.Println(p.IsValid()) // true/false

    // Or from string
    p2, _ := parser.NewParserFromString("WORKDIR /app\nCREATEFILE x/y.txt")
    program, err := p2.Parse()
    if err != nil {
        panic(err)
    }
    fmt.Println(program.ToString()) // JSON AST
}
```

## License

MIT — see [LICENSE](LICENSE).
