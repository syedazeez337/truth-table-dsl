package engine

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/syedazeez337/truthy/parser/grammar"
)

// Eval evaluates an Expr parse tree under env using a direct recursive descent over
// the generated context types. This avoids any visitor-dispatch pitfalls.
func Eval(tree antlr.ParseTree, env map[string]bool) bool {
	if env == nil {
		env = map[string]bool{}
	}
	ctx, ok := tree.(parser.IExprContext)
	if !ok {
		// Not an expr tree; be safe.
		return false
	}
	return evalExpr(ctx, env)
}

func evalExpr(ctx parser.IExprContext, env map[string]bool) bool {
	switch c := ctx.(type) {

	case *parser.NotContext:
		return !evalExpr(c.Expr(), env)

	case *parser.AndContext:
		return evalExpr(c.Expr(0), env) && evalExpr(c.Expr(1), env)

	case *parser.OrContext:
		return evalExpr(c.Expr(0), env) || evalExpr(c.Expr(1), env)

	case *parser.ParensContext:
		return evalExpr(c.Expr(), env)

	case *parser.TrueLitContext:
		return true

	case *parser.FalseLitContext:
		return false

	case *parser.VarContext:
		name := c.IDENT().GetText()
		return env[name] // defaults to false if unset

	default:
		// Unknown node type (shouldn't happen if grammar matches).
		return false
	}
}
