.PHONY:

# TODO build for mac and test it
clean:
	rm -rfv SnapEngine

linux: clean
	env CGO_ENABLED=1 CC=gcc GOOS=linux GOARCH=amd64 go build -o SnapEngine --race main.go
