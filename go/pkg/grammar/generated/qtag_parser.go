// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package generated // QTag
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type QTagParser struct {
	*antlr.BaseParser
}

var qtagParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func qtagParserInit() {
	staticData := &qtagParserStaticData
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
		"body", "comment", "lineComment", "multilineComment", "commentBody",
		"commentPrefix", "jsonBody",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 13, 75, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 5, 0, 16, 8, 0, 10, 0, 12, 0, 19, 9, 0,
		1, 0, 5, 0, 22, 8, 0, 10, 0, 12, 0, 25, 9, 0, 1, 0, 5, 0, 28, 8, 0, 10,
		0, 12, 0, 31, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 37, 8, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 5, 4, 48, 8, 4, 10, 4, 12, 4,
		51, 9, 4, 1, 4, 1, 4, 5, 4, 55, 8, 4, 10, 4, 12, 4, 58, 9, 4, 1, 4, 1,
		4, 1, 5, 5, 5, 63, 8, 5, 10, 5, 12, 5, 66, 9, 5, 1, 6, 1, 6, 5, 6, 70,
		8, 6, 10, 6, 12, 6, 73, 9, 6, 1, 6, 4, 17, 29, 64, 71, 0, 7, 0, 2, 4, 6,
		8, 10, 12, 0, 3, 2, 0, 5, 5, 8, 8, 1, 1, 9, 10, 1, 0, 12, 12, 75, 0, 23,
		1, 0, 0, 0, 2, 36, 1, 0, 0, 0, 4, 38, 1, 0, 0, 0, 6, 42, 1, 0, 0, 0, 8,
		49, 1, 0, 0, 0, 10, 64, 1, 0, 0, 0, 12, 67, 1, 0, 0, 0, 14, 16, 9, 0, 0,
		0, 15, 14, 1, 0, 0, 0, 16, 19, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0, 17, 15,
		1, 0, 0, 0, 18, 20, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 20, 22, 3, 2, 1, 0,
		21, 17, 1, 0, 0, 0, 22, 25, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 23, 24, 1,
		0, 0, 0, 24, 29, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 26, 28, 9, 0, 0, 0, 27,
		26, 1, 0, 0, 0, 28, 31, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 29, 27, 1, 0, 0,
		0, 30, 32, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 32, 33, 5, 0, 0, 1, 33, 1, 1,
		0, 0, 0, 34, 37, 3, 4, 2, 0, 35, 37, 3, 6, 3, 0, 36, 34, 1, 0, 0, 0, 36,
		35, 1, 0, 0, 0, 37, 3, 1, 0, 0, 0, 38, 39, 7, 0, 0, 0, 39, 40, 3, 8, 4,
		0, 40, 41, 7, 1, 0, 0, 41, 5, 1, 0, 0, 0, 42, 43, 5, 6, 0, 0, 43, 44, 3,
		8, 4, 0, 44, 45, 5, 7, 0, 0, 45, 7, 1, 0, 0, 0, 46, 48, 5, 11, 0, 0, 47,
		46, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0,
		0, 50, 52, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 52, 56, 3, 10, 5, 0, 53, 55,
		5, 11, 0, 0, 54, 53, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0,
		56, 57, 1, 0, 0, 0, 57, 59, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 59, 60, 3,
		12, 6, 0, 60, 9, 1, 0, 0, 0, 61, 63, 8, 2, 0, 0, 62, 61, 1, 0, 0, 0, 63,
		66, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 65, 11, 1, 0, 0,
		0, 66, 64, 1, 0, 0, 0, 67, 71, 5, 12, 0, 0, 68, 70, 9, 0, 0, 0, 69, 68,
		1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0,
		72, 13, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 8, 17, 23, 29, 36, 49, 56, 64,
		71,
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

// QTagParserInit initializes any static state used to implement QTagParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewQTagParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func QTagParserInit() {
	staticData := &qtagParserStaticData
	staticData.once.Do(qtagParserInit)
}

// NewQTagParser produces a new parser instance for the optional input antlr.TokenStream.
func NewQTagParser(input antlr.TokenStream) *QTagParser {
	QTagParserInit()
	this := new(QTagParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &qtagParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// QTagParser tokens.
const (
	QTagParserEOF            = antlr.TokenEOF
	QTagParserBT_IDENTIFIER  = 1
	QTagParserDQ_IDENTIFIER  = 2
	QTagParserSQ_LITERAL     = 3
	QTagParserDOLLAR_LITERAL = 4
	QTagParserDASH_COMMENT   = 5
	QTagParserOPENC          = 6
	QTagParserCLOSEC         = 7
	QTagParserC_COMMENT      = 8
	QTagParserNL             = 9
	QTagParserCR             = 10
	QTagParserSPACE          = 11
	QTagParserCURLY_O        = 12
	QTagParserOTHER          = 13
)

// QTagParser rules.
const (
	QTagParserRULE_body             = 0
	QTagParserRULE_comment          = 1
	QTagParserRULE_lineComment      = 2
	QTagParserRULE_multilineComment = 3
	QTagParserRULE_commentBody      = 4
	QTagParserRULE_commentPrefix    = 5
	QTagParserRULE_jsonBody         = 6
)

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QTagParserRULE_body
	return p
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QTagParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) EOF() antlr.TerminalNode {
	return s.GetToken(QTagParserEOF, 0)
}

func (s *BodyContext) AllComment() []ICommentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommentContext); ok {
			len++
		}
	}

	tst := make([]ICommentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommentContext); ok {
			tst[i] = t.(ICommentContext)
			i++
		}
	}

	return tst
}

func (s *BodyContext) Comment(i int) ICommentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentContext); ok {
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

	return t.(ICommentContext)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QTagListener); ok {
		listenerT.EnterBody(s)
	}
}

func (s *BodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QTagListener); ok {
		listenerT.ExitBody(s)
	}
}

func (p *QTagParser) Body() (localctx IBodyContext) {
	this := p
	_ = this

	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QTagParserRULE_body)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(17)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

			for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1+1 {
					p.SetState(14)
					p.MatchWildcard()

				}
				p.SetState(19)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
			}
			{
				p.SetState(20)
				p.Comment()
			}

		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(26)
			p.MatchWildcard()

		}
		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	{
		p.SetState(32)
		p.Match(QTagParserEOF)
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QTagParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QTagParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) LineComment() ILineCommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineCommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineCommentContext)
}

func (s *CommentContext) MultilineComment() IMultilineCommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultilineCommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultilineCommentContext)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QTagListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QTagListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *QTagParser) Comment() (localctx ICommentContext) {
	this := p
	_ = this

	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, QTagParserRULE_comment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(36)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case QTagParserDASH_COMMENT, QTagParserC_COMMENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.LineComment()
		}

	case QTagParserOPENC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.MultilineComment()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILineCommentContext is an interface to support dynamic dispatch.
type ILineCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineCommentContext differentiates from other interfaces.
	IsLineCommentContext()
}

type LineCommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineCommentContext() *LineCommentContext {
	var p = new(LineCommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QTagParserRULE_lineComment
	return p
}

func (*LineCommentContext) IsLineCommentContext() {}

func NewLineCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineCommentContext {
	var p = new(LineCommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QTagParserRULE_lineComment

	return p
}

func (s *LineCommentContext) GetParser() antlr.Parser { return s.parser }

func (s *LineCommentContext) CommentBody() ICommentBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentBodyContext)
}

func (s *LineCommentContext) C_COMMENT() antlr.TerminalNode {
	return s.GetToken(QTagParserC_COMMENT, 0)
}

func (s *LineCommentContext) DASH_COMMENT() antlr.TerminalNode {
	return s.GetToken(QTagParserDASH_COMMENT, 0)
}

func (s *LineCommentContext) NL() antlr.TerminalNode {
	return s.GetToken(QTagParserNL, 0)
}

func (s *LineCommentContext) CR() antlr.TerminalNode {
	return s.GetToken(QTagParserCR, 0)
}

func (s *LineCommentContext) EOF() antlr.TerminalNode {
	return s.GetToken(QTagParserEOF, 0)
}

func (s *LineCommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineCommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineCommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QTagListener); ok {
		listenerT.EnterLineComment(s)
	}
}

func (s *LineCommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QTagListener); ok {
		listenerT.ExitLineComment(s)
	}
}

func (p *QTagParser) LineComment() (localctx ILineCommentContext) {
	this := p
	_ = this

	localctx = NewLineCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, QTagParserRULE_lineComment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		_la = p.GetTokenStream().LA(1)

		if !(_la == QTagParserDASH_COMMENT || _la == QTagParserC_COMMENT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(39)
		p.CommentBody()
	}
	{
		p.SetState(40)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la - -1)) & ^0x3f) == 0 && ((int64(1)<<(_la - -1))&3073) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMultilineCommentContext is an interface to support dynamic dispatch.
type IMultilineCommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultilineCommentContext differentiates from other interfaces.
	IsMultilineCommentContext()
}

type MultilineCommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultilineCommentContext() *MultilineCommentContext {
	var p = new(MultilineCommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QTagParserRULE_multilineComment
	return p
}

func (*MultilineCommentContext) IsMultilineCommentContext() {}

func NewMultilineCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultilineCommentContext {
	var p = new(MultilineCommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QTagParserRULE_multilineComment

	return p
}

func (s *MultilineCommentContext) GetParser() antlr.Parser { return s.parser }

func (s *MultilineCommentContext) OPENC() antlr.TerminalNode {
	return s.GetToken(QTagParserOPENC, 0)
}

func (s *MultilineCommentContext) CommentBody() ICommentBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentBodyContext)
}

func (s *MultilineCommentContext) CLOSEC() antlr.TerminalNode {
	return s.GetToken(QTagParserCLOSEC, 0)
}

func (s *MultilineCommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultilineCommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultilineCommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QTagListener); ok {
		listenerT.EnterMultilineComment(s)
	}
}

func (s *MultilineCommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QTagListener); ok {
		listenerT.ExitMultilineComment(s)
	}
}

func (p *QTagParser) MultilineComment() (localctx IMultilineCommentContext) {
	this := p
	_ = this

	localctx = NewMultilineCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, QTagParserRULE_multilineComment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(QTagParserOPENC)
	}
	{
		p.SetState(43)
		p.CommentBody()
	}
	{
		p.SetState(44)
		p.Match(QTagParserCLOSEC)
	}

	return localctx
}

// ICommentBodyContext is an interface to support dynamic dispatch.
type ICommentBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPrefix returns the prefix rule contexts.
	GetPrefix() ICommentPrefixContext

	// GetJson returns the json rule contexts.
	GetJson() IJsonBodyContext

	// SetPrefix sets the prefix rule contexts.
	SetPrefix(ICommentPrefixContext)

	// SetJson sets the json rule contexts.
	SetJson(IJsonBodyContext)

	// IsCommentBodyContext differentiates from other interfaces.
	IsCommentBodyContext()
}

type CommentBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	prefix ICommentPrefixContext
	json   IJsonBodyContext
}

func NewEmptyCommentBodyContext() *CommentBodyContext {
	var p = new(CommentBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QTagParserRULE_commentBody
	return p
}

func (*CommentBodyContext) IsCommentBodyContext() {}

func NewCommentBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentBodyContext {
	var p = new(CommentBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QTagParserRULE_commentBody

	return p
}

func (s *CommentBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentBodyContext) GetPrefix() ICommentPrefixContext { return s.prefix }

func (s *CommentBodyContext) GetJson() IJsonBodyContext { return s.json }

func (s *CommentBodyContext) SetPrefix(v ICommentPrefixContext) { s.prefix = v }

func (s *CommentBodyContext) SetJson(v IJsonBodyContext) { s.json = v }

func (s *CommentBodyContext) CommentPrefix() ICommentPrefixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentPrefixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentPrefixContext)
}

func (s *CommentBodyContext) JsonBody() IJsonBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJsonBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJsonBodyContext)
}

func (s *CommentBodyContext) AllSPACE() []antlr.TerminalNode {
	return s.GetTokens(QTagParserSPACE)
}

func (s *CommentBodyContext) SPACE(i int) antlr.TerminalNode {
	return s.GetToken(QTagParserSPACE, i)
}

func (s *CommentBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QTagListener); ok {
		listenerT.EnterCommentBody(s)
	}
}

func (s *CommentBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QTagListener); ok {
		listenerT.ExitCommentBody(s)
	}
}

func (p *QTagParser) CommentBody() (localctx ICommentBodyContext) {
	this := p
	_ = this

	localctx = NewCommentBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, QTagParserRULE_commentBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(46)
				p.Match(QTagParserSPACE)
			}

		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	{
		p.SetState(52)

		var _x = p.CommentPrefix()

		localctx.(*CommentBodyContext).prefix = _x
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == QTagParserSPACE {
		{
			p.SetState(53)
			p.Match(QTagParserSPACE)
		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(59)

		var _x = p.JsonBody()

		localctx.(*CommentBodyContext).json = _x
	}

	return localctx
}

// ICommentPrefixContext is an interface to support dynamic dispatch.
type ICommentPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentPrefixContext differentiates from other interfaces.
	IsCommentPrefixContext()
}

type CommentPrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentPrefixContext() *CommentPrefixContext {
	var p = new(CommentPrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QTagParserRULE_commentPrefix
	return p
}

func (*CommentPrefixContext) IsCommentPrefixContext() {}

func NewCommentPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentPrefixContext {
	var p = new(CommentPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QTagParserRULE_commentPrefix

	return p
}

func (s *CommentPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentPrefixContext) AllCURLY_O() []antlr.TerminalNode {
	return s.GetTokens(QTagParserCURLY_O)
}

func (s *CommentPrefixContext) CURLY_O(i int) antlr.TerminalNode {
	return s.GetToken(QTagParserCURLY_O, i)
}

func (s *CommentPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QTagListener); ok {
		listenerT.EnterCommentPrefix(s)
	}
}

func (s *CommentPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QTagListener); ok {
		listenerT.ExitCommentPrefix(s)
	}
}

func (p *QTagParser) CommentPrefix() (localctx ICommentPrefixContext) {
	this := p
	_ = this

	localctx = NewCommentPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, QTagParserRULE_commentPrefix)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(61)
				_la = p.GetTokenStream().LA(1)

				if _la <= 0 || _la == QTagParserCURLY_O {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IJsonBodyContext is an interface to support dynamic dispatch.
type IJsonBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonBodyContext differentiates from other interfaces.
	IsJsonBodyContext()
}

type JsonBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonBodyContext() *JsonBodyContext {
	var p = new(JsonBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QTagParserRULE_jsonBody
	return p
}

func (*JsonBodyContext) IsJsonBodyContext() {}

func NewJsonBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonBodyContext {
	var p = new(JsonBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QTagParserRULE_jsonBody

	return p
}

func (s *JsonBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonBodyContext) CURLY_O() antlr.TerminalNode {
	return s.GetToken(QTagParserCURLY_O, 0)
}

func (s *JsonBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QTagListener); ok {
		listenerT.EnterJsonBody(s)
	}
}

func (s *JsonBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QTagListener); ok {
		listenerT.ExitJsonBody(s)
	}
}

func (p *QTagParser) JsonBody() (localctx IJsonBodyContext) {
	this := p
	_ = this

	localctx = NewJsonBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, QTagParserRULE_jsonBody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(QTagParserCURLY_O)
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			p.SetState(68)
			p.MatchWildcard()

		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}
