// Code generated from ./grammar/TruthExpr.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TruthExpr
import "github.com/antlr4-go/antlr/v4"

// BaseTruthExprListener is a complete listener for a parse tree produced by TruthExprParser.
type BaseTruthExprListener struct{}

var _ TruthExprListener = &BaseTruthExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTruthExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTruthExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTruthExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTruthExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNot is called when production Not is entered.
func (s *BaseTruthExprListener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production Not is exited.
func (s *BaseTruthExprListener) ExitNot(ctx *NotContext) {}

// EnterOr is called when production Or is entered.
func (s *BaseTruthExprListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production Or is exited.
func (s *BaseTruthExprListener) ExitOr(ctx *OrContext) {}

// EnterParens is called when production Parens is entered.
func (s *BaseTruthExprListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BaseTruthExprListener) ExitParens(ctx *ParensContext) {}

// EnterVar is called when production Var is entered.
func (s *BaseTruthExprListener) EnterVar(ctx *VarContext) {}

// ExitVar is called when production Var is exited.
func (s *BaseTruthExprListener) ExitVar(ctx *VarContext) {}

// EnterAnd is called when production And is entered.
func (s *BaseTruthExprListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production And is exited.
func (s *BaseTruthExprListener) ExitAnd(ctx *AndContext) {}

// EnterFalseLit is called when production FalseLit is entered.
func (s *BaseTruthExprListener) EnterFalseLit(ctx *FalseLitContext) {}

// ExitFalseLit is called when production FalseLit is exited.
func (s *BaseTruthExprListener) ExitFalseLit(ctx *FalseLitContext) {}

// EnterTrueLit is called when production TrueLit is entered.
func (s *BaseTruthExprListener) EnterTrueLit(ctx *TrueLitContext) {}

// ExitTrueLit is called when production TrueLit is exited.
func (s *BaseTruthExprListener) ExitTrueLit(ctx *TrueLitContext) {}
