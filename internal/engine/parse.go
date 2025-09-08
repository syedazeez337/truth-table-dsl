package engine

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/syedazeez337/truthy/parser/grammar"
)

/******** error counting listener ********/
type errCounter struct{ *antlr.DefaultErrorListener; N int }

func (e *errCounter) SyntaxError(_ antlr.Recognizer, _ interface{}, _ int, _ int, _ string, _ antlr.RecognitionException) {
	e.N++
}

/******** identifier collector (Listener) ********/
type identCollector struct {
	*parser.BaseTruthExprListener
	Set map[string]struct{}
}

func (l *identCollector) EnterVar(ctx *parser.VarContext) {
	name := ctx.IDENT().GetText()
	l.Set[name] = struct{}{}
}

/******** ParseOne: returns parse tree, identifier set, error count ********/
func ParseOne(expr string) (antlr.ParseTree, map[string]struct{}, int) {
	ec := &errCounter{DefaultErrorListener: &antlr.DefaultErrorListener{}}

	input := antlr.NewInputStream(expr)

	// LEXER with error listener
	lex := parser.NewTruthExprLexer(input)
	lex.RemoveErrorListeners()
	lex.AddErrorListener(ec)

	tokens := antlr.NewCommonTokenStream(lex, 0)

	// PARSER with error listener
	p := parser.NewTruthExprParser(tokens)
	p.BuildParseTrees = true
	p.RemoveErrorListeners()
	p.AddErrorListener(ec)

	tree := p.Expr()

	// ensure EOF only (no trailing junk)
	if tokens.LA(1) != antlr.TokenEOF {
		ec.N++
	}

	// collect identifiers
	set := map[string]struct{}{}
	collector := &identCollector{Set: set}
	antlr.ParseTreeWalkerDefault.Walk(collector, tree)

	return tree, set, ec.N
}
