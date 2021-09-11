# GoSH
A simple implementation of a UNIX shell.

## Requirements
You need Golang version >= 1.16 installed on your system. You can install one for your platform [here](https://golang.org/dl/).

```
$ go version
go version go1.16.5 darwin/amd64
```

## A Simple Example
```
$ go run .

[username@host:/path/to/current/dir] $ ls -al | grep git
drwxr-xr-x  12 user  group   384  9 Jul 11:05 .git
-rw-r--r--   1 user  group    21  2 Jul 09:00 .gitignore
```

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
=== RUN   TestPipedCommand
--- PASS: TestPipedCommand (0.00s)
=== RUN   TestOutputRedirection
--- PASS: TestOutputRedirection (0.00s)
=== RUN   TestInputRedirection
--- PASS: TestInputRedirection (0.00s)
=== RUN   TestAppendOutputRedirection
--- PASS: TestAppendOutputRedirection (0.00s)
=== RUN   TestRedirectStdoutStderrToSameFile
--- PASS: TestRedirectStdoutStderrToSameFile (0.00s)
=== RUN   TestAppendStdoutStderrToSameFile
--- PASS: TestAppendStdoutStderrToSameFile (0.00s)
PASS
coverage: 97.4% of statements
ok      github.com/anthonyabeo/gosh/parser      0.952s  coverage: 97.4% of statements

To see your test coverage in the browser
$ go tool cover -html=coverage.out
```
