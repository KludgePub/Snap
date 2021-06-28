.PHONY:

# TODO build for mac and test it
clean:
	rm -rfv Snap

debian: clean
	env CGO_ENABLED=1 CC=gcc GOOS=linux GOARCH=amd64 go build -tags -gcflags="all=-N -l" -o Snap --race main.go
