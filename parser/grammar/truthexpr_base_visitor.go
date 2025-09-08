// Code generated from ./grammar/TruthExpr.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TruthExpr
import "github.com/antlr4-go/antlr/v4"

type BaseTruthExprVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTruthExprVisitor) VisitNot(ctx *NotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTruthExprVisitor) VisitOr(ctx *OrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTruthExprVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTruthExprVisitor) VisitVar(ctx *VarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTruthExprVisitor) VisitAnd(ctx *AndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTruthExprVisitor) VisitFalseLit(ctx *FalseLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTruthExprVisitor) VisitTrueLit(ctx *TrueLitContext) interface{} {
	return v.VisitChildren(ctx)
}
