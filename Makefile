BINARY_NAME=gdiff

all: build

build:
	go build -o $(BINARY_NAME) *.go

clean:
	rm -f $(BINARY_NAME)

.PHONY: all build clean