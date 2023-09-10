package hedi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokens_String(t *testing.T) {
	tokens := Tokens{{Value: "a"}, {Value: "b"}}
	assert.Equal(t, "ab", tokens.String())
}

func TestTokens_JSON(t *testing.T) {
	tokens := Tokens{{Type: SegmentIdentifier, Value: "ISA"}, {Type: ElementDelimiter, Value: "*"}}
	assert.Equal(t, `[{"type":"segment_identifier","value":"ISA"},{"type":"element_delimiter","value":"*"}]`, string(tokens.JSON()))

	tokens = Tokens{}
	assert.Equal(t, `[]`, string(tokens.JSON()))
}
