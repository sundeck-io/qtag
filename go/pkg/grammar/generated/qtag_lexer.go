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
		"", "", "", "", "", "'qtag:'", "'--'", "'/*'", "'*/'", "'//'", "'\\n'",
		"'\\r'", "", "'{'",
	}
	staticData.symbolicNames = []string{
		"", "BT_IDENTIFIER", "DQ_IDENTIFIER", "SQ_LITERAL", "DOLLAR_LITERAL",
		"QTAG_PREFIX", "DASH_COMMENT", "OPENC", "CLOSEC", "C_COMMENT", "NL",
		"CR", "SPACE", "CURLY_O", "OTHER",
	}
	staticData.ruleNames = []string{
		"BT_IDENTIFIER", "DQ_IDENTIFIER", "SQ_LITERAL", "DOLLAR_LITERAL", "QTAG_PREFIX",
		"DASH_COMMENT", "OPENC", "CLOSEC", "C_COMMENT", "NL", "CR", "SPACE",
		"CURLY_O", "OTHER",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 14, 108, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0,
		1, 0, 5, 0, 34, 8, 0, 10, 0, 12, 0, 37, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 5, 1, 45, 8, 1, 10, 1, 12, 1, 48, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 58, 8, 2, 10, 2, 12, 2, 61, 9, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 71, 8, 3, 10, 3, 12, 3, 74,
		9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 4, 13, 105, 8, 13, 11, 13, 12,
		13, 106, 1, 106, 0, 14, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 1, 0, 4, 1, 0, 96, 96,
		1, 0, 34, 34, 2, 0, 39, 39, 92, 92, 2, 0, 9, 9, 32, 32, 117, 0, 1, 1, 0,
		0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0,
		0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1,
		0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25,
		1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 1, 29, 1, 0, 0, 0, 3, 40, 1, 0, 0, 0, 5,
		51, 1, 0, 0, 0, 7, 64, 1, 0, 0, 0, 9, 77, 1, 0, 0, 0, 11, 83, 1, 0, 0,
		0, 13, 86, 1, 0, 0, 0, 15, 89, 1, 0, 0, 0, 17, 92, 1, 0, 0, 0, 19, 95,
		1, 0, 0, 0, 21, 97, 1, 0, 0, 0, 23, 99, 1, 0, 0, 0, 25, 101, 1, 0, 0, 0,
		27, 104, 1, 0, 0, 0, 29, 35, 5, 96, 0, 0, 30, 34, 8, 0, 0, 0, 31, 32, 5,
		96, 0, 0, 32, 34, 5, 96, 0, 0, 33, 30, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0,
		34, 37, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 38, 1,
		0, 0, 0, 37, 35, 1, 0, 0, 0, 38, 39, 5, 96, 0, 0, 39, 2, 1, 0, 0, 0, 40,
		46, 5, 34, 0, 0, 41, 45, 8, 1, 0, 0, 42, 43, 5, 34, 0, 0, 43, 45, 5, 34,
		0, 0, 44, 41, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46, 44,
		1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 49, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0,
		49, 50, 5, 34, 0, 0, 50, 4, 1, 0, 0, 0, 51, 59, 5, 39, 0, 0, 52, 53, 5,
		92, 0, 0, 53, 58, 9, 0, 0, 0, 54, 55, 5, 39, 0, 0, 55, 58, 5, 39, 0, 0,
		56, 58, 8, 2, 0, 0, 57, 52, 1, 0, 0, 0, 57, 54, 1, 0, 0, 0, 57, 56, 1,
		0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60,
		62, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 62, 63, 5, 39, 0, 0, 63, 6, 1, 0, 0,
		0, 64, 65, 5, 36, 0, 0, 65, 66, 5, 36, 0, 0, 66, 72, 1, 0, 0, 0, 67, 71,
		8, 1, 0, 0, 68, 69, 5, 34, 0, 0, 69, 71, 5, 34, 0, 0, 70, 67, 1, 0, 0,
		0, 70, 68, 1, 0, 0, 0, 71, 74, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 72, 73,
		1, 0, 0, 0, 73, 75, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 75, 76, 5, 34, 0, 0,
		76, 8, 1, 0, 0, 0, 77, 78, 5, 113, 0, 0, 78, 79, 5, 116, 0, 0, 79, 80,
		5, 97, 0, 0, 80, 81, 5, 103, 0, 0, 81, 82, 5, 58, 0, 0, 82, 10, 1, 0, 0,
		0, 83, 84, 5, 45, 0, 0, 84, 85, 5, 45, 0, 0, 85, 12, 1, 0, 0, 0, 86, 87,
		5, 47, 0, 0, 87, 88, 5, 42, 0, 0, 88, 14, 1, 0, 0, 0, 89, 90, 5, 42, 0,
		0, 90, 91, 5, 47, 0, 0, 91, 16, 1, 0, 0, 0, 92, 93, 5, 47, 0, 0, 93, 94,
		5, 47, 0, 0, 94, 18, 1, 0, 0, 0, 95, 96, 5, 10, 0, 0, 96, 20, 1, 0, 0,
		0, 97, 98, 5, 13, 0, 0, 98, 22, 1, 0, 0, 0, 99, 100, 7, 3, 0, 0, 100, 24,
		1, 0, 0, 0, 101, 102, 5, 123, 0, 0, 102, 26, 1, 0, 0, 0, 103, 105, 9, 0,
		0, 0, 104, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0,
		106, 104, 1, 0, 0, 0, 107, 28, 1, 0, 0, 0, 10, 0, 33, 35, 44, 46, 57, 59,
		70, 72, 106, 0,
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
	QTagLexerQTAG_PREFIX    = 5
	QTagLexerDASH_COMMENT   = 6
	QTagLexerOPENC          = 7
	QTagLexerCLOSEC         = 8
	QTagLexerC_COMMENT      = 9
	QTagLexerNL             = 10
	QTagLexerCR             = 11
	QTagLexerSPACE          = 12
	QTagLexerCURLY_O        = 13
	QTagLexerOTHER          = 14
)
