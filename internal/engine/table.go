package engine

import (
	"fmt"
	"sort"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

/******** Truth table helpers ********/

func allAssignments(vars []string) []map[string]bool {
	n := len(vars)
	if n == 0 {
		return []map[string]bool{{}}
	}
	total := 1 << n
	rows := make([]map[string]bool, 0, total)
	for mask := 0; mask < total; mask++ {
		env := make(map[string]bool, n)
		for i, name := range vars {
			bit := (mask >> (n - 1 - i)) & 1 // MSB changes slowest
			env[name] = bit == 1
		}
		rows = append(rows, env)
	}
	return rows
}

func formatHeader(vars []string) string {
	if len(vars) == 0 {
		return "expr"
	}
	return strings.Join(append(vars, "expr"), "\t")
}

func formatRow(vars []string, env map[string]bool, result bool) string {
	out := make([]string, 0, len(vars)+1)
	for _, v := range vars {
		if env[v] {
			out = append(out, "T")
		} else {
			out = append(out, "F")
		}
	}
	if result {
		out = append(out, "T")
	} else {
		out = append(out, "F")
	}
	return strings.Join(out, "\t")
}

func BuildTruthTable(tree antlr.ParseTree, varsSet map[string]struct{}) []string {
	var vars []string
	for k := range varsSet {
		vars = append(vars, k)
	}
	sort.Strings(vars)

	lines := []string{formatHeader(vars)}
	for _, env := range allAssignments(vars) {
		val := Eval(tree, env)
		lines = append(lines, formatRow(vars, env, val))
	}
	return lines
}

func PrintTruthTable(tree antlr.ParseTree, varsSet map[string]struct{}) {
	for _, line := range BuildTruthTable(tree, varsSet) {
		fmt.Println(line)
	}
}
