# GoSH
A simple implementation of the UNIX shell.

## Requirements
You need Golang version 1.16 installed on your system. You can install one for your platform [here](https://golang.org/dl/).


## Testing
```
$ go test -v ./parser -cover -coverprofile=coverage.out
=== RUN   TestTokens
--- PASS: TestTokens (0.00s)
=== RUN   TestParser
--- PASS: TestParser (0.00s)
=== RUN   TestParseCmd
--- PASS: TestParseCmd (0.00s)
=== RUN   TestParseCommand
--- PASS: TestParseCommand (0.00s)
PASS
coverage: 94.4% of statements
ok      github.com/anthonyabeo/gosh/parser      0.760s  coverage: 94.4% of statements

To see your test coverage in the browser
$ go tool cover -html=coverage.out
```