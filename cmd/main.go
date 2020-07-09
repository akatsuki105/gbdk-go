package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
)

var version string

var (
	showVersion = flag.Bool("v", false, "show version")
	output      = flag.String("o", "game.gb", "output file path")
	help        = flag.Bool("help", false, "show detailed help")
)

func Run() int {
	os.RemoveAll("./tmp")
	flag.Parse()
	dir := flag.Arg(0)

	if *showVersion {
		printVersion()
		return ExitCodeOK
	}

	if *help {
		printHelp()
		return ExitCodeOK
	}

	assets := filepath.Join(dir, "asset", "*.go")
	assetCmd := fmt.Sprintf("./go2c %s", assets)
	fmt.Println(assetCmd)
	cmd := exec.Command("zsh", "-c", assetCmd)
	cmd.Env = append(os.Environ(), "ASSET=true")
	cmd.Output()

	scripts := filepath.Join(dir, "*.go")
	scriptCmd := fmt.Sprintf("./go2c %s", scripts)
	fmt.Println(scriptCmd)
	if _, err := exec.Command("zsh", "-c", scriptCmd).Output(); err != nil {
		fmt.Fprintf(os.Stderr, "compile error: %s", err)
		return ExitCodeError
	}

	zshCmd := fmt.Sprintf("bin/lcc -o %s tmp/*.c", *output)
	fmt.Println(zshCmd)
	if _, err := exec.Command("zsh", "-c", zshCmd).Output(); err != nil {
		fmt.Fprintf(os.Stderr, "build error: %s", err)
		return ExitCodeError
	}

	os.RemoveAll("./tmp")
	fmt.Println("build success!")
	return ExitCodeOK
}

func main() {
	os.Exit(Run())
}

func printVersion() {
	if version == "" {
		version = "develop"
	}
	fmt.Println(version)
}

func printHelp() {
	if version == "" {
		version = "develop"
	}
	fmt.Printf(`gbdk-go: %s

Usage
  $ gbdkgo [options] dir

Options:
  -help
        show detailed help
  -o string
        output gb file (default "game.gb")
  -v    show version
`, version)
}
