ifdef COMSPEC
	EXE_EXT := .exe
else
	EXE_EXT := 
endif

.PHONY: build
build:
	go build -o gbdkgo -ldflags "-X main.version=$(shell git describe --tags)" ./cmd/