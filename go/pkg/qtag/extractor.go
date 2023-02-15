package qtag

import (
	_ "embed"
	"encoding/json"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	grammar "github.com/sundeck-io/CommentQL/go/pkg/grammar/generated"
	"gopkg.in/yaml.v3"
	"strings"
)

type treeShapeListener struct {
	*grammar.BaseQTagListener
	comments []jsonComment
}

//go:embed declared.yaml
var declarations string

func (s *treeShapeListener) EnterCommentBody(ctx *grammar.CommentBodyContext) {
	prefix := ctx.GetPrefix().GetText()
	jsonText := ctx.GetJson().GetText()
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonText), &jsonMap)
	if err != nil {
		return
	}
	s.comments = append(s.comments, jsonComment{prefix, jsonMap})
}

type jsonComment struct {
	prefix   string
	metadata map[string]interface{}
}

type Extractor interface {
	Extract(query string, includeUnidentifiedComments bool, includeUnknownAttributes bool) []QTag
}

type extractor struct {
	Extractor
	declarations []Declaration
}

func NewExtractor() Extractor {
	var decls []Declaration
	err := yaml.Unmarshal([]byte(declarations), &decls)
	if err != nil {
		panic(err)
	}

	return extractor{declarations: decls}
}

// Extract finds all query tags within the provided query string
func (e extractor) Extract(query string, includeUnidentifiedComments bool, includeUnknownAttributes bool) []QTag {
	input := antlr.NewInputStream(query)
	lexer := grammar.NewQTagLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := grammar.NewQTagParser(stream)
	p.Interpreter.SetPredictionMode(antlr.PredictionModeSLL)

	//TODO: replace error listener
	p.AddErrorListener(antlr.NewConsoleErrorListener())
	p.BuildParseTrees = true
	tree := p.Body()
	listener := &treeShapeListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	var allTags []QTag

	for _, comment := range listener.comments {

		// TODO: Avoid the linear traversal here. This is just v1 for simplicity.
		for _, declaration := range e.declarations {
			tags := declaration.extract(comment, includeUnknownAttributes)

			if tags != nil {
				allTags = append(allTags, tags...)
				break
			}
		}

		if includeUnidentifiedComments {
			// create a generic element.
		}

	}
	return allTags
}

type QTag struct {
	Source string
	Type   Type
	Key    string
	Value  interface{}
}

type Type string

const (
	TRACE       Type = "TRACE"
	USER        Type = "USER"
	DIMENSION   Type = "DIMENSION"
	MEASURE     Type = "MEASURE"
	APPLICATION Type = "APPLICATION"
	UNDECLARED  Type = "UNDECLARED"
)

func (s Type) String() string {
	switch s {
	case TRACE:
		return "TRACE"
	case USER:
		return "USER"
	case DIMENSION:
		return "DIMENSION"
	case MEASURE:
		return "MEASURE"
	case APPLICATION:
		return "APPLICATION"
	case UNDECLARED:
		return "UNDECLARED"
	}
	return "unknown"
}

type Declaration struct {
	Identifier Identifier `yaml:identifier`
	Name       string     `yaml:name`
	Site       string     `yaml:site`
	Fields     []Field    `yaml:fields`
	fm         *map[string]Field
}

func (d Declaration) findAttribute(name string) (Field, bool) {
	if d.fm == nil {
		fieldMap := make(map[string]Field)
		for _, field := range d.Fields {
			fieldMap[field.Name] = field
		}
		d.fm = &fieldMap
	}
	val, ok := (*d.fm)[name]
	return val, ok
}

// extract the attributes from the comment if they match the expected identifier
func (d Declaration) extract(comment jsonComment, includeUnknownAttributes bool) []QTag {
	if d.Identifier.Prefix != "" {
		if !strings.EqualFold(d.Identifier.Prefix, comment.prefix) {
			return nil
		}
	}

	if d.Identifier.Fields != nil {
		for k, v := range d.Identifier.Fields {
			if v != comment.metadata[k] {
				return nil
			}
		}
	}

	// we matched by identifier.
	var tags []QTag
	for k, v := range comment.metadata {
		val, found := d.findAttribute(k)
		if found {
			tags = append(tags, QTag{
				Source: d.Name,
				Type:   val.Type,
				Key:    k,
				Value:  v,
			})
		} else if includeUnknownAttributes {
			tags = append(tags, QTag{
				Source: d.Name,
				Type:   UNDECLARED,
				Key:    k,
				Value:  v,
			})
		}
	}

	if len(tags) == 0 {
		return nil
	}

	return tags
}

type Identifier struct {
	Prefix string
	Fields map[string]interface{}
}

type Field struct {
	Name string
	Type Type
}

func AsJson(tags []QTag) string {
	str, err := json.MarshalIndent(tags, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(str)
}
