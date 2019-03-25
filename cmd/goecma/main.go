package main

import (
	"flag"
	"fmt"
	"github.com/keizo042/goecma/js"
	"os"
)

var (
	Version string
)

func usage() {
	usage := []string{
		fmt.Sprintf("version: %s", Version),
		"goecma [option] [file]",
	}
	for _, u := range usage {
		fmt.Println(u)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		usage()
		return
	}
	path := flag.Args()[0]
	if path == "" {
		fmt.Println("option required")
		return
	}
	state, err := js.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	tokens, err := state.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	ast, err := state.Parse(tokens)
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := state.Run(ast, 0, nil); err != nil {
		fmt.Println(err)
	}
}
