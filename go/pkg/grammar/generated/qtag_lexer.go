// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package generated

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type QTagLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var qtaglexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func qtaglexerLexerInit() {
	staticData := &qtaglexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "", "", "", "", "'--'", "'/*'", "'*/'", "'//'", "'\\n'", "'\\r'",
		"", "'{'",
	}
	staticData.symbolicNames = []string{
		"", "BT_IDENTIFIER", "DQ_IDENTIFIER", "SQ_LITERAL", "DOLLAR_LITERAL",
		"DASH_COMMENT", "OPENC", "CLOSEC", "C_COMMENT", "NL", "CR", "SPACE",
		"CURLY_O", "OTHER",
	}
	staticData.ruleNames = []string{
		"BT_IDENTIFIER", "DQ_IDENTIFIER", "SQ_LITERAL", "DOLLAR_LITERAL", "DASH_COMMENT",
		"OPENC", "CLOSEC", "C_COMMENT", "NL", "CR", "SPACE", "CURLY_O", "OTHER",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 13, 100, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 32,
		8, 0, 10, 0, 12, 0, 35, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1,
		43, 8, 1, 10, 1, 12, 1, 46, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 5, 2, 56, 8, 2, 10, 2, 12, 2, 59, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 69, 8, 3, 10, 3, 12, 3, 72, 9, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 4, 12, 97,
		8, 12, 11, 12, 12, 12, 98, 1, 98, 0, 13, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 1, 0, 4, 1,
		0, 96, 96, 1, 0, 34, 34, 2, 0, 39, 39, 92, 92, 2, 0, 9, 9, 32, 32, 109,
		0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 1, 27, 1, 0, 0, 0, 3, 38, 1, 0, 0, 0, 5, 49, 1,
		0, 0, 0, 7, 62, 1, 0, 0, 0, 9, 75, 1, 0, 0, 0, 11, 78, 1, 0, 0, 0, 13,
		81, 1, 0, 0, 0, 15, 84, 1, 0, 0, 0, 17, 87, 1, 0, 0, 0, 19, 89, 1, 0, 0,
		0, 21, 91, 1, 0, 0, 0, 23, 93, 1, 0, 0, 0, 25, 96, 1, 0, 0, 0, 27, 33,
		5, 96, 0, 0, 28, 32, 8, 0, 0, 0, 29, 30, 5, 96, 0, 0, 30, 32, 5, 96, 0,
		0, 31, 28, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 32, 35, 1, 0, 0, 0, 33, 31,
		1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 36, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0,
		36, 37, 5, 96, 0, 0, 37, 2, 1, 0, 0, 0, 38, 44, 5, 34, 0, 0, 39, 43, 8,
		1, 0, 0, 40, 41, 5, 34, 0, 0, 41, 43, 5, 34, 0, 0, 42, 39, 1, 0, 0, 0,
		42, 40, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 44, 45, 1,
		0, 0, 0, 45, 47, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 47, 48, 5, 34, 0, 0, 48,
		4, 1, 0, 0, 0, 49, 57, 5, 39, 0, 0, 50, 51, 5, 92, 0, 0, 51, 56, 9, 0,
		0, 0, 52, 53, 5, 39, 0, 0, 53, 56, 5, 39, 0, 0, 54, 56, 8, 2, 0, 0, 55,
		50, 1, 0, 0, 0, 55, 52, 1, 0, 0, 0, 55, 54, 1, 0, 0, 0, 56, 59, 1, 0, 0,
		0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 57,
		1, 0, 0, 0, 60, 61, 5, 39, 0, 0, 61, 6, 1, 0, 0, 0, 62, 63, 5, 36, 0, 0,
		63, 64, 5, 36, 0, 0, 64, 70, 1, 0, 0, 0, 65, 69, 8, 1, 0, 0, 66, 67, 5,
		34, 0, 0, 67, 69, 5, 34, 0, 0, 68, 65, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0,
		69, 72, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 73, 1,
		0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 74, 5, 34, 0, 0, 74, 8, 1, 0, 0, 0, 75,
		76, 5, 45, 0, 0, 76, 77, 5, 45, 0, 0, 77, 10, 1, 0, 0, 0, 78, 79, 5, 47,
		0, 0, 79, 80, 5, 42, 0, 0, 80, 12, 1, 0, 0, 0, 81, 82, 5, 42, 0, 0, 82,
		83, 5, 47, 0, 0, 83, 14, 1, 0, 0, 0, 84, 85, 5, 47, 0, 0, 85, 86, 5, 47,
		0, 0, 86, 16, 1, 0, 0, 0, 87, 88, 5, 10, 0, 0, 88, 18, 1, 0, 0, 0, 89,
		90, 5, 13, 0, 0, 90, 20, 1, 0, 0, 0, 91, 92, 7, 3, 0, 0, 92, 22, 1, 0,
		0, 0, 93, 94, 5, 123, 0, 0, 94, 24, 1, 0, 0, 0, 95, 97, 9, 0, 0, 0, 96,
		95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 98, 96, 1, 0, 0,
		0, 99, 26, 1, 0, 0, 0, 10, 0, 31, 33, 42, 44, 55, 57, 68, 70, 98, 0,
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

// QTagLexerInit initializes any static state used to implement QTagLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewQTagLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func QTagLexerInit() {
	staticData := &qtaglexerLexerStaticData
	staticData.once.Do(qtaglexerLexerInit)
}

// NewQTagLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewQTagLexer(input antlr.CharStream) *QTagLexer {
	QTagLexerInit()
	l := new(QTagLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &qtaglexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "QTag.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// QTagLexer tokens.
const (
	QTagLexerBT_IDENTIFIER  = 1
	QTagLexerDQ_IDENTIFIER  = 2
	QTagLexerSQ_LITERAL     = 3
	QTagLexerDOLLAR_LITERAL = 4
	QTagLexerDASH_COMMENT   = 5
	QTagLexerOPENC          = 6
	QTagLexerCLOSEC         = 7
	QTagLexerC_COMMENT      = 8
	QTagLexerNL             = 9
	QTagLexerCR             = 10
	QTagLexerSPACE          = 11
	QTagLexerCURLY_O        = 12
	QTagLexerOTHER          = 13
)
