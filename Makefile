BINNAME = go-exchange-chart-fetcher
GO = go
PKGS = `glide nv`

glide:
	glide update ./...

test:
	export CWD=${PWD} && $(GO) test -cover $(PKGS)

test-verbose:
	export CWD=${PWD} && $(GO) test -cover $(PKGS)

test-short:
	export CWD=${PWD} && $(GO) test -cover $(PKGS) -short

build:
	$(GO) build $(PKGS)

run:
	$(GO) run *.go config.yml

.PHONY: glide test test-verbose build run
