BINARY_NAME=beerctl

build:
	GOARCH=amd64 GOOS=linux go build -o $(BINARY_NAME) .

install: build
	sudo ln -s $(shell pwd)/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
	. ~/.bashrc

cleanup:
	go clean
	rm $(BINARY_NAME)
	sudo rm /usr/local/bin/$(BINARY_NAME)
