# gbdk-go

GameBoy Development Kit for gopher.

gbdk-go is a Go binding for GBDK. You can do GameBoy software development with Go.

## Architecture

<img src="./architecture.svg">

gbdk-go compiles Go source code and outputs C code.

The output C code is built into GB ROM by GBDK.

```go
package main

import "github.com/Akatsuki-py/gbdk-go/api/stdio"

func main() {
	stdio.Printf("Hello World!")
	stdio.Printf("\n\nPress Start")
}
```

The above Go code is compiled into the following C code.

```c
#include <stdio.h>
void main() {
    printf("Hello World!");
    printf("\n\nPress Start");
}
```

## Usage

```sh
$ gbdkgo [options] dir
```

#### Example

```sh
$ gbdkgo example/simple_shmup
```

## API

Detailed documents is ToDo.

The API is basically a GBDK binding.

Some are not yet implemented.

## Project rules

Detailed documents is ToDo.

The build target is the Go directory, not the file. 

The directory structure also needs to be specific.

## Linter

Although gbdk-go allows GB development with Go, there are some grammars such as defer and goroutine that cannot be used. 

So I'm going to write a linter for gbdk-go.
