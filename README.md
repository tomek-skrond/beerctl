# beerctl

`beerctl` is a CLI tool written in Go that fetches data from [Punk API](https://punkapi.com/documentation/v2).

### Setup

To set up this tool, you can use `make` tool:

```
$ make build
```

Above command builds a binary with the name specified in `Makefile`
```
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
```

After running `build`, you can access tool:

```
$ ./beerctl <options>
```

If you want to use this tool globally in the system, after running `make build`, run `make install`. This command sets up a symbolic link in your `/usr/local/bin/` folder.

To clean up, enter `make cleanup` command.

### Usage

Available options: `-all`, `-id <int>`, `-random`

```
$ beerctl -help

Usage of beerctl:
  -all
    	View all available beers
  -id int
    	View beer with specified ID (default -1)
  -random
    	View a randomly selected beer
```
