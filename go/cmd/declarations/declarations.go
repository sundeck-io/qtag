// generate.go
package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/goccy/go-yaml"
)

// Field represents a field in the YAML file
type Field struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

// Config represents the structure of the YAML file
type Config struct {
	Fields     []Field        `yaml:"fields"`
	Name       string         `yaml:"name"`
	Identifier map[string]any `yaml:"identifier"`
}

// MethodData represents data needed to generate a method
type MethodData struct {
	ClassName  string
	MethodName string
	Name       string
}

// ClassData represents data needed to generate a class
type ClassData struct {
	ClassName  string
	Name       string
	Identifier string
}

const structTemplate = `
type {{.ClassName}} struct {
	Builder
	name string
    identifier map[string]any
    values map[string]any
}

func New{{.ClassName}}() *{{.ClassName}} {
	x := {{.ClassName}}{}
	x.init()
	return &x
}

func (c *{{.ClassName}}) init() {
c.name = "{{.Name}}"
c.identifier= make(map[string]any)
json.Unmarshal([]byte("{{.Identifier}}"), &c.identifier)
c.values = make(map[string]any)
}

func (c *{{.ClassName}}) Format() (string, error) {
	return format(c.name, c.identifier, c.values)
}

func (c *{{.ClassName}}) UnknownValue(name string, value any) {
c.values[name] = value
}
`

// Template for generating methods
const methodTemplate = `

func (c *{{.ClassName}}) {{.MethodName}}(value any) {
    c.values["{{.Name}}"] = value
}
`

func CreateFile(path string) (*strings.Builder, error) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	b, _ := json.Marshal(config.Identifier)

	tmplM, err := template.New("method").Parse(methodTemplate)
	if err != nil {
		return nil, err
	}
	tmplC, err := template.New("class").Parse(structTemplate)
	if err != nil {
		return nil, err
	}

	c := strings.Builder{}
	escapedIdentifier := strings.ReplaceAll(string(b), `"`, `\"`)
	cData := ClassData{
		ClassName:  capitalize(config.Name),
		Name:       config.Name,
		Identifier: escapedIdentifier,
	}
	err = tmplC.Execute(&c, cData)
	if err != nil {
		return nil, err
	}

	// Iterate through fields and generate methods
	for _, field := range config.Fields {
		data := MethodData{
			ClassName:  capitalize(config.Name),
			MethodName: "Add" + capitalize(removeDash(field.Name)),
			Name:       removeDash(field.Name),
		}
		err = tmplM.Execute(&c, data)
		if err != nil {
			return nil, err
		}
	}
	return &c, nil
}

func main() {
	// Read and parse the YAML file
	yamlDir := "pkg/qtag/declarations"
	files, err := ioutil.ReadDir(yamlDir)
	if err != nil {
		panic(err)
	}

	c := strings.Builder{}
	c.WriteString("package qtag\nimport (\n\t\"encoding/json\"\n)\n\n")
	for _, f := range files {
		var s *strings.Builder
		s, err = CreateFile(filepath.Join(yamlDir, f.Name()))
		if err != nil {
			panic(err)
		}
		c.WriteString(s.String())
	}

	// Create the generated file
	f, err := os.Create("pkg/qtag/generated.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	formattedCode, err := format.Source([]byte(c.String()))
	if err != nil {
		panic(err)
	}
	f.Write(formattedCode)
	fmt.Println("example/generated.go has been generated")
}

// capitalize capitalizes the first letter of a string
func capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	return string(s[0]-32) + s[1:]
}

func removeDash(s string) string {
	if len(s) == 0 {
		return ""
	}
	if strings.Contains(s, "-") {
		return strings.Replace(s, "-", "__", -1)
	}
	return s
}
