package dotenv

import (
	"testing"

	"gotest.tools/v3/assert"
)

var testInput = `
a=b
a[1]=c
a.propertyKey=d
árvíztűrő-TÜKÖRFÚRÓGÉP=ÁRVÍZTŰRŐ-tükörfúrógép
`

func TestParseBytes(t *testing.T) {
	p := newParser()

	expectedOutput := map[string]string{
		"a":                      "b",
		"a[1]":                   "c",
		"a.propertyKey":          "d",
		"árvíztűrő-TÜKÖRFÚRÓGÉP": "ÁRVÍZTŰRŐ-tükörfúrógép",
	}

	out := map[string]string{}
	err := p.parse(testInput, out, nil)

	assert.NilError(t, err)
	assert.Equal(t, len(expectedOutput), len(out))
	for key, value := range expectedOutput {
		assert.Equal(t, value, out[key])
	}
}
