// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package generated // QTag
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// QTagListener is a complete listener for a parse tree produced by QTagParser.
type QTagListener interface {
	antlr.ParseTreeListener

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterLineComment is called when entering the lineComment production.
	EnterLineComment(c *LineCommentContext)

	// EnterMultilineComment is called when entering the multilineComment production.
	EnterMultilineComment(c *MultilineCommentContext)

	// EnterCommentBody is called when entering the commentBody production.
	EnterCommentBody(c *CommentBodyContext)

	// EnterCommentPrefix is called when entering the commentPrefix production.
	EnterCommentPrefix(c *CommentPrefixContext)

	// EnterJsonBody is called when entering the jsonBody production.
	EnterJsonBody(c *JsonBodyContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitLineComment is called when exiting the lineComment production.
	ExitLineComment(c *LineCommentContext)

	// ExitMultilineComment is called when exiting the multilineComment production.
	ExitMultilineComment(c *MultilineCommentContext)

	// ExitCommentBody is called when exiting the commentBody production.
	ExitCommentBody(c *CommentBodyContext)

	// ExitCommentPrefix is called when exiting the commentPrefix production.
	ExitCommentPrefix(c *CommentPrefixContext)

	// ExitJsonBody is called when exiting the jsonBody production.
	ExitJsonBody(c *JsonBodyContext)
}
