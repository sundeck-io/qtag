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
