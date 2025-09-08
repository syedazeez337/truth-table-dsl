// Code generated from ./grammar/TruthExpr.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TruthExpr
import "github.com/antlr4-go/antlr/v4"

// TruthExprListener is a complete listener for a parse tree produced by TruthExprParser.
type TruthExprListener interface {
	antlr.ParseTreeListener

	// EnterNot is called when entering the Not production.
	EnterNot(c *NotContext)

	// EnterOr is called when entering the Or production.
	EnterOr(c *OrContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterVar is called when entering the Var production.
	EnterVar(c *VarContext)

	// EnterAnd is called when entering the And production.
	EnterAnd(c *AndContext)

	// EnterFalseLit is called when entering the FalseLit production.
	EnterFalseLit(c *FalseLitContext)

	// EnterTrueLit is called when entering the TrueLit production.
	EnterTrueLit(c *TrueLitContext)

	// ExitNot is called when exiting the Not production.
	ExitNot(c *NotContext)

	// ExitOr is called when exiting the Or production.
	ExitOr(c *OrContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitVar is called when exiting the Var production.
	ExitVar(c *VarContext)

	// ExitAnd is called when exiting the And production.
	ExitAnd(c *AndContext)

	// ExitFalseLit is called when exiting the FalseLit production.
	ExitFalseLit(c *FalseLitContext)

	// ExitTrueLit is called when exiting the TrueLit production.
	ExitTrueLit(c *TrueLitContext)
}
