package qtag

import (
	"github.com/goccy/go-yaml"
	"gotest.tools/v3/assert"
	"os"
	"sort"
	"testing"
)

type QueryFixture struct {
	Name          string
	Query         string
	All           bool
	AllAttributes bool
	Outcome       []QTag
}

func Test(t *testing.T) {
	yamlFile, err := os.ReadFile("../../../tests/test.yaml")
	if err != nil {
		panic(err)
	}

	var fixtures []QueryFixture
	err = yaml.UnmarshalWithOptions(yamlFile, &fixtures, yaml.Strict())
	if err != nil {
		panic(err)
	}

	e := NewExtractor()
	for _, fixture := range fixtures {
		t.Run(fixture.Name, func(t *testing.T) {
			tags := e.Extract(fixture.Query, fixture.All, fixture.AllAttributes)
			sortTags(tags)
			sortTags(fixture.Outcome)
			assert.DeepEqual(t, fixture.Outcome, tags)
		})
	}
}

func sortTags(tags []QTag) {
	sort.Slice(tags[:], func(i, j int) bool {
		if tags[i].Source != tags[j].Source {
			return tags[i].Source < tags[j].Source
		}

		if tags[i].Key != tags[j].Key {
			return tags[i].Key < tags[j].Key
		}

		// TODO: add value sort for consistency.
		//if tags[i].Value != tags[j].Value {
		//	return tags[i].Value < tags[j].Value
		//}

		return tags[i].Type < tags[j].Type
	})
}
