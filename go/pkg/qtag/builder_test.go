package qtag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	var x = NewDbt()
	x.AddConnection_name("x")
	s, err := x.Format()
	assert.NoError(t, err)
	assert.Equal(t, `{"app":"dbt","connection_name":"x"}`, s)
}

func TestMerge(t *testing.T) {
	var x = NewDbt()
	x.AddConnection_name("x")
	x.AddNode_id("bob")
	var y = NewDbt()
	y.AddConnection_name("xx")
	y.AddIs_incremental(false)
	x.Merge(y)
	s, err := x.Format()
	assert.NoError(t, err)
	assert.Equal(t, `{"app":"dbt","connection_name":"xx","is_incremental":false,"node_id":"bob"}`, s)
}
