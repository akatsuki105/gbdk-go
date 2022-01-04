package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
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
	assetCmd := fmt.Sprintf("go2c %s", assets)
	fmt.Println(assetCmd)
	cmd := exec.Command("/bin/sh", "-c", assetCmd)
	cmd.Env = append(os.Environ(), "ASSET=true")
	cmd.Output()

	scripts := filepath.Join(dir, "*.go")
	scriptCmd := fmt.Sprintf("go2c %s", scripts)
	fmt.Println(scriptCmd)
	if out, err := exec.Command("/bin/sh", "-c", scriptCmd).Output(); err != nil {
		fmt.Fprintf(os.Stderr, "compile error: %s\n%s\n", err, out)
		return ExitCodeError
	}

	zshCmd := fmt.Sprintf("%s -o %s tmp/*.c", path.Join(gbdkPath(), "lcc"), *output)

	fmt.Println(zshCmd)
	if out, err := exec.Command("/bin/sh", "-c", zshCmd).Output(); err != nil {
		fmt.Fprintf(os.Stderr, "\nbuild error: %s\n%s\n", err, out)
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

func gbdkPath() string {
	if p := os.Getenv("GBDKDIR"); p != "" {
		return path.Join(p, "bin")
	}

	return ""
}
