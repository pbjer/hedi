package hedi

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLexISA(t *testing.T) {
	t.Run("Line feed as segment terminator should succeed", func(t *testing.T) {
		file, err := os.Open("./test/850_with_new_line_segment_terminator.txt")
		assert.NoError(t, err)
		defer file.Close()

		tokens, separators, err := lexISA(file)
		assert.NoError(t, err)
		assert.Equal(t, 33, len(tokens))
		assert.Equal(t, int32(10), separators.Segment)
		assert.Equal(t, int32(42), separators.Element)
		assert.Equal(t, int32(62), separators.SubElement)
	})
	t.Run("Tilde as segment terminator should succeed", func(t *testing.T) {
		file, err := os.Open("./test/850_with_tilde_segment_terminator.txt")
		assert.NoError(t, err)
		defer file.Close()

		tokens, separators, err := lexISA(file)
		assert.NoError(t, err)
		assert.Equal(t, 33, len(tokens))
		assert.Equal(t, int32(126), separators.Segment)
		assert.Equal(t, int32(42), separators.Element)
		assert.Equal(t, int32(62), separators.SubElement)
	})
}

func TestLex(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		file, err := os.Open("./test/850_with_tilde_segment_terminator.txt")
		defer file.Close()
		assert.NoError(t, err)
		lexer := NewLexer(file)
		tokens, err := lexer.Tokens()
		assert.NoError(t, err)
		assert.Len(t, tokens, 503)
	})
}
