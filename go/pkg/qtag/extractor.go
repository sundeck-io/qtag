package qtag

import (
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/goccy/go-yaml"
	grammar "github.com/sundeck-io/CommentQL/go/pkg/grammar/generated"
	"strings"
)

type treeShapeListener struct {
	*grammar.BaseQTagListener
	comments []jsonComment
}

//go:embed declarations/*.yaml
var declarationFolder embed.FS

func (s *treeShapeListener) EnterCommentBody(ctx *grammar.CommentBodyContext) {
	prefix := ctx.GetPrefix().GetText()
	jsonText := ctx.GetJson().GetText()
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonText), &jsonMap)
	if err != nil {
		return
	}
	s.comments = append(s.comments, jsonComment{prefix: prefix, metadata: jsonMap, body: jsonText})
}

type jsonComment struct {
	prefix   string
	metadata map[string]interface{}
	body     string
}

type Extractor interface {
	Extract(query string, arbitraryComments bool, arbitraryAttributes bool) []QTag
}

type extractor struct {
	Extractor
	declarations []Declaration
}

func NewExtractor() Extractor {
	var decls []Declaration
	items, err := declarationFolder.ReadDir("declarations")
	if err != nil {
		panic(err)
	}

	for _, r := range items {
		name := "declarations/" + r.Name()
		if r.IsDir() {
			continue
		}
		data, err := declarationFolder.ReadFile(name)
		if err != nil {
			panic(fmt.Errorf("failure reading file %v. %w", r.Name(), err))
		}
		var declaration Declaration
		err = yaml.UnmarshalWithOptions(data, &declaration, yaml.Strict())
		if err != nil {
			panic(fmt.Errorf("failure parsing yaml %v. %w", r.Name(), err))
		}
		decls = append(decls, declaration)
	}
	return extractor{declarations: decls}
}

// Extract finds all query tags within the provided query string
func (e extractor) Extract(query string, arbitraryComments bool, arbitraryAttributes bool) []QTag {
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
		foundDeclaration := false

		isSelfDescribing := len(comment.prefix) >= 5 && strings.EqualFold(comment.prefix[0:4], "qtag")

		if !isSelfDescribing {
			// TODO: Avoid the linear traversal here. This is just v1 for simplicity.
			for _, declaration := range e.declarations {
				tags := declaration.extract(comment, arbitraryComments || arbitraryAttributes)

				if tags != nil {
					foundDeclaration = true
					allTags = append(allTags, tags...)

				}
			}

			if foundDeclaration || !arbitraryComments {
				continue
			}

		}

		// is selfdescribing or arbitrary
		var tags []QTag
		source := "other"
		if isSelfDescribing {
			splitSource := strings.SplitN(comment.prefix, ":", 2)
			if len(splitSource) == 2 {
				source = splitSource[1]
			}
		} else {
			if comment.prefix != "" {
				source = comment.prefix
			}
		}

		for k, v := range comment.metadata {
			keyName := k
			valueType := UNKNOWN

			if isSelfDescribing {
				splitKey := strings.SplitN(k, ":", 2)

				if len(splitKey) == 2 {
					keyName = splitKey[0]
					valueType = asType(splitKey[1])
				}
			}

			tags = append(tags, QTag{
				Source: source,
				Type:   valueType,
				Key:    keyName,
				Value:  v,
			})
		}

		allTags = append(allTags, tags...)

	}
	return allTags
}

type QTag struct {
	Source string      `json:SOURCE`
	Key    string      `json:KEY`
	Value  interface{} `json:VALUE`
	Type   Type        `json:TYPE`
}

type Type string

const (
	TRACE       Type = "TRACE"
	USER        Type = "USER"
	DIMENSION   Type = "DIMENSION"
	MEASURE     Type = "MEASURE"
	APPLICATION Type = "APPLICATION"
	UNDECLARED  Type = "UNDECLARED"
	UNKNOWN     Type = "UNKNOWN"
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

func asType(s string) Type {
	switch strings.ToUpper(s) {
	case "TRACE":
		return TRACE
	case "USER":
		return USER
	case "DIMENSION":
		return DIMENSION
	case "MEASURE":
		return MEASURE
	case "APPLICATION":
		return APPLICATION
	case "UNDECLARED":
		return UNDECLARED
	}
	return UNKNOWN
}

type Declaration struct {
	Identifier identifier `yaml:identifier`
	Name       string     `yaml:name`
	Site       string     `yaml:site`
	Fields     []field    `yaml:fields`
	fm         *map[string]field
}

func (d Declaration) findAttribute(name string) (field, bool) {
	if d.fm == nil {
		fieldMap := make(map[string]field)
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

	if d.Identifier.BodyMatch != "" {
		if !strings.Contains(comment.body, d.Identifier.BodyMatch) {
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

type identifier struct {
	Prefix    string
	Fields    map[string]interface{}
	BodyMatch string
}

type field struct {
	Name string
	Type Type
}
