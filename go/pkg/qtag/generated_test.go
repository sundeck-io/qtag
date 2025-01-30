package qtag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSundeck_AddIcecap(t *testing.T) {
	i := &Icecap{}
	i.init()
	i.AddTarget_size("Small")
	i.AddNo_route_reason([]string{"MissingScalarFunction", "UnableToResolve"})
	i.AddUsed(true)

	c := &Sundeck{}
	c.init()
	c.AddIcecap(i.Values())
	tags, err := c.Format()
	require.NoError(t, err)
	expectedTags := `{"app":"sundeck","icecap":{"no_route_reason":["MissingScalarFunction","UnableToResolve"],"target_size":"Small","used":true}}`
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
			require.Equal(t, "Small", icMap["target_size"])
			require.Equal(t, []interface{}{"MissingScalarFunction", "UnableToResolve"}, icMap["no_route_reason"])
			require.Equal(t, true, icMap["used"])
		} else {
			require.Fail(t, "unexpected tag")
		}
	}
}
