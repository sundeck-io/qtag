// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package generated // QTag
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseQTagListener is a complete listener for a parse tree produced by QTagParser.
type BaseQTagListener struct{}

var _ QTagListener = &BaseQTagListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQTagListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQTagListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQTagListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQTagListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterBody is called when production body is entered.
func (s *BaseQTagListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BaseQTagListener) ExitBody(ctx *BodyContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseQTagListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseQTagListener) ExitComment(ctx *CommentContext) {}

// EnterLineComment is called when production lineComment is entered.
func (s *BaseQTagListener) EnterLineComment(ctx *LineCommentContext) {}

// ExitLineComment is called when production lineComment is exited.
func (s *BaseQTagListener) ExitLineComment(ctx *LineCommentContext) {}

// EnterMultilineComment is called when production multilineComment is entered.
func (s *BaseQTagListener) EnterMultilineComment(ctx *MultilineCommentContext) {}

// ExitMultilineComment is called when production multilineComment is exited.
func (s *BaseQTagListener) ExitMultilineComment(ctx *MultilineCommentContext) {}

// EnterCommentBody is called when production commentBody is entered.
func (s *BaseQTagListener) EnterCommentBody(ctx *CommentBodyContext) {}

// ExitCommentBody is called when production commentBody is exited.
func (s *BaseQTagListener) ExitCommentBody(ctx *CommentBodyContext) {}

// EnterCommentPrefix is called when production commentPrefix is entered.
func (s *BaseQTagListener) EnterCommentPrefix(ctx *CommentPrefixContext) {}

// ExitCommentPrefix is called when production commentPrefix is exited.
func (s *BaseQTagListener) ExitCommentPrefix(ctx *CommentPrefixContext) {}

// EnterJsonBody is called when production jsonBody is entered.
func (s *BaseQTagListener) EnterJsonBody(ctx *JsonBodyContext) {}

// ExitJsonBody is called when production jsonBody is exited.
func (s *BaseQTagListener) ExitJsonBody(ctx *JsonBodyContext) {}
