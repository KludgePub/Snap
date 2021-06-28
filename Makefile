.PHONY:

# TODO build for mac and test it
clean:
	rm -rfv SnapEngine

debian: clean
	env CGO_ENABLED=1 CC=gcc GOOS=linux GOARCH=amd64 go build -tags -gcflags="all=-N -l" -o SnapEngine --race main.go
