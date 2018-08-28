GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=passphrase
BINARY_NAME_WINDOWS=passphrase.exe

.PHONY: all build build-windows test clean

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

build-windows:
	GOOS=windows \
	GOARCH=386 \
	$(GOBUILD) -o $(BINARY_NAME_WINDOWS) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME_WINDOWS)