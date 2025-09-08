package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/syedazeez337/truthy/internal/engine"
)

func readExpr() (string, error) {
	if len(os.Args) > 1 {
		return strings.Join(os.Args[1:], " "), nil
	}
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	expr, err := readExpr()
	if err != nil {
		fmt.Fprintln(os.Stderr, "read:", err)
		os.Exit(1)
	}

	tree, vars, errs := engine.ParseOne(expr)
	if errs > 0 {
		fmt.Fprintln(os.Stderr, "parse: syntax errors found")
		os.Exit(2)
	}
	engine.PrintTruthTable(tree, vars)
}
