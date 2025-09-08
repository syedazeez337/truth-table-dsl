// Code generated from ./grammar/TruthExpr.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type TruthExprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var TruthExprLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func truthexprlexerLexerInit() {
	staticData := &TruthExprLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'!'", "'&'", "'|'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "TRUE", "FALSE", "IDENT", "WS", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "TRUE", "FALSE", "IDENT", "WS",
		"LINE_COMMENT", "ESC",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 10, 72, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 5, 7, 47, 8, 7, 10, 7, 12, 7, 50, 9, 7, 1, 8, 4, 8, 53, 8, 8, 11,
		8, 12, 8, 54, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 63, 8, 9, 10, 9,
		12, 9, 66, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 0, 0, 11, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 0, 1, 0, 12,
		2, 0, 84, 84, 116, 116, 2, 0, 82, 82, 114, 114, 2, 0, 85, 85, 117, 117,
		2, 0, 69, 69, 101, 101, 2, 0, 70, 70, 102, 102, 2, 0, 65, 65, 97, 97, 2,
		0, 76, 76, 108, 108, 2, 0, 83, 83, 115, 115, 3, 0, 65, 90, 95, 95, 97,
		122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32,
		2, 0, 10, 10, 13, 13, 73, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 1,
		23, 1, 0, 0, 0, 3, 25, 1, 0, 0, 0, 5, 27, 1, 0, 0, 0, 7, 29, 1, 0, 0, 0,
		9, 31, 1, 0, 0, 0, 11, 33, 1, 0, 0, 0, 13, 38, 1, 0, 0, 0, 15, 44, 1, 0,
		0, 0, 17, 52, 1, 0, 0, 0, 19, 58, 1, 0, 0, 0, 21, 69, 1, 0, 0, 0, 23, 24,
		5, 33, 0, 0, 24, 2, 1, 0, 0, 0, 25, 26, 5, 38, 0, 0, 26, 4, 1, 0, 0, 0,
		27, 28, 5, 124, 0, 0, 28, 6, 1, 0, 0, 0, 29, 30, 5, 40, 0, 0, 30, 8, 1,
		0, 0, 0, 31, 32, 5, 41, 0, 0, 32, 10, 1, 0, 0, 0, 33, 34, 7, 0, 0, 0, 34,
		35, 7, 1, 0, 0, 35, 36, 7, 2, 0, 0, 36, 37, 7, 3, 0, 0, 37, 12, 1, 0, 0,
		0, 38, 39, 7, 4, 0, 0, 39, 40, 7, 5, 0, 0, 40, 41, 7, 6, 0, 0, 41, 42,
		7, 7, 0, 0, 42, 43, 7, 3, 0, 0, 43, 14, 1, 0, 0, 0, 44, 48, 7, 8, 0, 0,
		45, 47, 7, 9, 0, 0, 46, 45, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 1,
		0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 16, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 51,
		53, 7, 10, 0, 0, 52, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 52, 1, 0,
		0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57, 6, 8, 0, 0, 57, 18,
		1, 0, 0, 0, 58, 59, 5, 47, 0, 0, 59, 60, 5, 47, 0, 0, 60, 64, 1, 0, 0,
		0, 61, 63, 8, 11, 0, 0, 62, 61, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64, 62,
		1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 67, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0,
		67, 68, 6, 9, 0, 0, 68, 20, 1, 0, 0, 0, 69, 70, 5, 92, 0, 0, 70, 71, 9,
		0, 0, 0, 71, 22, 1, 0, 0, 0, 4, 0, 48, 54, 64, 1, 6, 0, 0,
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

// TruthExprLexerInit initializes any static state used to implement TruthExprLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewTruthExprLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TruthExprLexerInit() {
	staticData := &TruthExprLexerLexerStaticData
	staticData.once.Do(truthexprlexerLexerInit)
}

// NewTruthExprLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewTruthExprLexer(input antlr.CharStream) *TruthExprLexer {
	TruthExprLexerInit()
	l := new(TruthExprLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &TruthExprLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "TruthExpr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TruthExprLexer tokens.
const (
	TruthExprLexerT__0         = 1
	TruthExprLexerT__1         = 2
	TruthExprLexerT__2         = 3
	TruthExprLexerT__3         = 4
	TruthExprLexerT__4         = 5
	TruthExprLexerTRUE         = 6
	TruthExprLexerFALSE        = 7
	TruthExprLexerIDENT        = 8
	TruthExprLexerWS           = 9
	TruthExprLexerLINE_COMMENT = 10
)
