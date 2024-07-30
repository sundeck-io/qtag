package qtag

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Builder interface {
	init(name string)
}

func format(name string, identifier map[string]any, values map[string]any) (string, error) {
	modifiedValues := make(map[string]any)
	for k, v := range values {
		if strings.Contains(k, "__") {
			modifiedValues[strings.ReplaceAll(k, "__", "-")] = v
		} else {
			modifiedValues[k] = v
		}
	}
	if f, ok := identifier["fields"]; ok {
		if m, ok := f.(map[string]any); ok {
			for k, v := range m {
				modifiedValues[k] = v
			}
			b, err := json.Marshal(modifiedValues)
			if err != nil {
				return "", err
			}
			return string(b), nil
		}
		return "", errors.New("field 'fields' should be a map")
	}
	if f, ok := identifier["prefix"]; ok {
		b, err := json.Marshal(modifiedValues)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s %s", f, string(b)), nil
	}
	return "", fmt.Errorf("unknown qtag format")
}

type UnknownQtag struct {
	Builder
	name       string
	identifier map[string]any
	values     map[string]any
}

func NewUnknownQtag(name string) *UnknownQtag {
	x := UnknownQtag{}
	x.init()
	x.name = name
	return &x
}

func (c *UnknownQtag) init() {
	c.identifier = make(map[string]any)
	json.Unmarshal([]byte("{\"fields\":{\"app\":\""+c.name+"\"}}"), &c.identifier)
	c.values = make(map[string]any)
}

func (c *UnknownQtag) Format() (string, error) {
	return format(c.name, c.identifier, c.values)
}

func (c *UnknownQtag) UnknownValue(name string, value any) {
	c.values[name] = value
}
