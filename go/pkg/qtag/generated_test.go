package qtag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSundeck_AddIcecap(t *testing.T) {

	m1 := make(map[string]any)
	m1["size"] = "Small"
	m1["reason"] = []string{"MissingFunction", "UnableToResolve"}
	m1["used"] = true

	c := &Sundeck{}
	c.init()
	c.AddIcecap(m1)
	tags, err := c.Format()
	require.NoError(t, err)
	expectedTags := `{"app":"sundeck","icecap":{"reason":["MissingFunction","UnableToResolve"],"size":"Small","used":true}}`
	require.Equal(t, expectedTags, tags)

	queryText := "select 1; -- " + tags
	extractor := NewExtractor()
	qtags := extractor.Extract(queryText, false, false)
	require.Len(t, qtags, 2)
	for _, tag := range qtags {
		if tag.Key == "app" {
			require.Equal(t, "sundeck", tag.Value)
		} else if tag.Key == "icecap" {
			icMap, ok := tag.Value.(map[string]interface{})
			require.True(t, ok)
			require.Equal(t, "Small", icMap["size"])
			require.Equal(t, []interface{}{"MissingFunction", "UnableToResolve"}, icMap["reason"])
			require.Equal(t, true, icMap["used"])
		} else {
			require.Fail(t, "unexpected tag")
		}
	}
}
