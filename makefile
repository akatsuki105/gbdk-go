ifdef COMSPEC
	EXE_EXT := .exe
else
	EXE_EXT := 
endif

.PHONY: build
build:
	go build -o go2c ./compiler/
	go build -o gbdkgo -ldflags "-X main.version=$(shell git describe --tags)" ./cmd/

.PHONY: compiler
compiler:
	go build -o go2c ./compiler/

.PHONY: clean
clean:
	-rm -f gbdkgo go2c game.gb