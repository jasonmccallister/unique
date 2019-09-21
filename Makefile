build:
	GOOS=darwin GOARCH=amd64 go build -o unique-macos .
	GOOS=windows GOARCH=386 go build -o unique-windows .
	GOOS=linux GOARCH=amd64 go build -o unique-linux .
