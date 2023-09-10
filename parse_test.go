package hedi

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseSegments(t *testing.T) {
	file, err := os.Open("./test/850_with_tilde_segment_terminator.txt")
	assert.NoError(t, err)
	parser := NewParser(file)
	segments, err := parser.Parse()
	assert.NoError(t, err)
	assert.Len(t, segments, 37)
}
