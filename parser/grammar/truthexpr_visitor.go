// Code generated from ./grammar/TruthExpr.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TruthExpr
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by TruthExprParser.
type TruthExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TruthExprParser#Not.
	VisitNot(ctx *NotContext) interface{}

	// Visit a parse tree produced by TruthExprParser#Or.
	VisitOr(ctx *OrContext) interface{}

	// Visit a parse tree produced by TruthExprParser#Parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by TruthExprParser#Var.
	VisitVar(ctx *VarContext) interface{}

	// Visit a parse tree produced by TruthExprParser#And.
	VisitAnd(ctx *AndContext) interface{}

	// Visit a parse tree produced by TruthExprParser#FalseLit.
	VisitFalseLit(ctx *FalseLitContext) interface{}

	// Visit a parse tree produced by TruthExprParser#TrueLit.
	VisitTrueLit(ctx *TrueLitContext) interface{}
}
