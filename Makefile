NAME:=$(shell basename `git rev-parse --show-toplevel`)
RELEASE:=$(shell git rev-parse --verify --short HEAD)
VERSION = 0.1.0

all: setbin

setbin: build
	cp repd /usr/local/bin

build:
	go build -ldflags "-X main.version=$(VERSION)-$(RELEASE)" -o repd

lint:
	go list ./... | grep -v /vendor/ | xargs -L1 golint
