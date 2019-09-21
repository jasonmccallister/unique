build:
	GOOS=darwin GOARCH=amd64 go build -o unique-darwin-amd64 .
	GOOS=windows GOARCH=386 go build -o unique-windows-amd64.exe .
	GOOS=linux GOARCH=amd64 go build -o unique-linux-amd64 .
