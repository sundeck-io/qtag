package qtag

import (
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func Test(t *testing.T) {
	e := NewExtractor()
	tags := e.Extract(`select * from foo
  -- Sigma Σ {"kind":"adhoc","request-id":"e0bd080279a9e63","email":"robert@sundeck.io"}`, true, true)

	assert.Equal(t, 3, len(tags))

	fmt.Println(AsJson(tags))

}
