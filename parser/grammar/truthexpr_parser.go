// Code generated from ./grammar/TruthExpr.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TruthExpr
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TruthExprParser struct {
	*antlr.BaseParser
}

var TruthExprParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func truthexprParserInit() {
	staticData := &TruthExprParserStaticData
	staticData.LiteralNames = []string{
		"", "'!'", "'&'", "'|'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "TRUE", "FALSE", "IDENT", "WS", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 10, 26, 2, 0, 7, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 3, 0, 13, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0,
		21, 8, 0, 10, 0, 12, 0, 24, 9, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 30, 0, 12,
		1, 0, 0, 0, 2, 3, 6, 0, -1, 0, 3, 4, 5, 1, 0, 0, 4, 13, 3, 0, 0, 7, 5,
		6, 5, 4, 0, 0, 6, 7, 3, 0, 0, 0, 7, 8, 5, 5, 0, 0, 8, 13, 1, 0, 0, 0, 9,
		13, 5, 6, 0, 0, 10, 13, 5, 7, 0, 0, 11, 13, 5, 8, 0, 0, 12, 2, 1, 0, 0,
		0, 12, 5, 1, 0, 0, 0, 12, 9, 1, 0, 0, 0, 12, 10, 1, 0, 0, 0, 12, 11, 1,
		0, 0, 0, 13, 22, 1, 0, 0, 0, 14, 15, 10, 6, 0, 0, 15, 16, 5, 2, 0, 0, 16,
		21, 3, 0, 0, 7, 17, 18, 10, 5, 0, 0, 18, 19, 5, 3, 0, 0, 19, 21, 3, 0,
		0, 6, 20, 14, 1, 0, 0, 0, 20, 17, 1, 0, 0, 0, 21, 24, 1, 0, 0, 0, 22, 20,
		1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 1, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0,
		3, 12, 20, 22,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// TruthExprParserInit initializes any static state used to implement TruthExprParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTruthExprParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TruthExprParserInit() {
	staticData := &TruthExprParserStaticData
	staticData.once.Do(truthexprParserInit)
}

// NewTruthExprParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTruthExprParser(input antlr.TokenStream) *TruthExprParser {
	TruthExprParserInit()
	this := new(TruthExprParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TruthExprParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "TruthExpr.g4"

	return this
}

// TruthExprParser tokens.
const (
	TruthExprParserEOF          = antlr.TokenEOF
	TruthExprParserT__0         = 1
	TruthExprParserT__1         = 2
	TruthExprParserT__2         = 3
	TruthExprParserT__3         = 4
	TruthExprParserT__4         = 5
	TruthExprParserTRUE         = 6
	TruthExprParserFALSE        = 7
	TruthExprParserIDENT        = 8
	TruthExprParserWS           = 9
	TruthExprParserLINE_COMMENT = 10
)

// TruthExprParserRULE_expr is the TruthExprParser rule.
const TruthExprParserRULE_expr = 0

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TruthExprParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TruthExprParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TruthExprParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NotContext struct {
	ExprContext
}

func NewNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotContext {
	var p = new(NotContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TruthExprListener); ok {
		listenerT.EnterNot(s)
	}
}

func (s *NotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TruthExprListener); ok {
		listenerT.ExitNot(s)
	}
}

func (s *NotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TruthExprVisitor:
		return t.VisitNot(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrContext struct {
	ExprContext
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OrContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TruthExprListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TruthExprListener); ok {
		listenerT.ExitOr(s)
	}
}

func (s *OrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TruthExprVisitor:
		return t.VisitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensContext struct {
	ExprContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TruthExprListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TruthExprListener); ok {
		listenerT.ExitParens(s)
	}
}

func (s *ParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TruthExprVisitor:
		return t.VisitParens(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarContext struct {
	ExprContext
}

func NewVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarContext {
	var p = new(VarContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *VarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarContext) IDENT() antlr.TerminalNode {
	return s.GetToken(TruthExprParserIDENT, 0)
}

func (s *VarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TruthExprListener); ok {
		listenerT.EnterVar(s)
	}
}

func (s *VarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TruthExprListener); ok {
		listenerT.ExitVar(s)
	}
}

func (s *VarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TruthExprVisitor:
		return t.VisitVar(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndContext struct {
	ExprContext
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AndContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TruthExprListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TruthExprListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (s *AndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TruthExprVisitor:
		return t.VisitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type FalseLitContext struct {
	ExprContext
}

func NewFalseLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseLitContext {
	var p = new(FalseLitContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FalseLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseLitContext) FALSE() antlr.TerminalNode {
	return s.GetToken(TruthExprParserFALSE, 0)
}

func (s *FalseLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TruthExprListener); ok {
		listenerT.EnterFalseLit(s)
	}
}

func (s *FalseLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TruthExprListener); ok {
		listenerT.ExitFalseLit(s)
	}
}

func (s *FalseLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TruthExprVisitor:
		return t.VisitFalseLit(s)

	default:
		return t.VisitChildren(s)
	}
}

type TrueLitContext struct {
	ExprContext
}

func NewTrueLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueLitContext {
	var p = new(TrueLitContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *TrueLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueLitContext) TRUE() antlr.TerminalNode {
	return s.GetToken(TruthExprParserTRUE, 0)
}

func (s *TrueLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TruthExprListener); ok {
		listenerT.EnterTrueLit(s)
	}
}

func (s *TrueLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TruthExprListener); ok {
		listenerT.ExitTrueLit(s)
	}
}

func (s *TrueLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TruthExprVisitor:
		return t.VisitTrueLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TruthExprParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *TruthExprParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, TruthExprParserRULE_expr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(12)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TruthExprParserT__0:
		localctx = NewNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(3)
			p.Match(TruthExprParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(4)
			p.expr(7)
		}

	case TruthExprParserT__3:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(5)
			p.Match(TruthExprParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(6)
			p.expr(0)
		}
		{
			p.SetState(7)
			p.Match(TruthExprParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TruthExprParserTRUE:
		localctx = NewTrueLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(9)
			p.Match(TruthExprParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TruthExprParserFALSE:
		localctx = NewFalseLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(10)
			p.Match(TruthExprParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TruthExprParserIDENT:
		localctx = NewVarContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(11)
			p.Match(TruthExprParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(20)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TruthExprParserRULE_expr)
				p.SetState(14)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(15)
					p.Match(TruthExprParserT__1)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(16)
					p.expr(7)
				}

			case 2:
				localctx = NewOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TruthExprParserRULE_expr)
				p.SetState(17)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(18)
					p.Match(TruthExprParserT__2)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(19)
					p.expr(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *TruthExprParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TruthExprParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
