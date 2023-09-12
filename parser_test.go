package hedi

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestParseSegments(t *testing.T) {
	file, err := os.Open("./test/850_with_tilde_segment_terminator.txt")
	assert.NoError(t, err)
	parser := NewParser(file)
	segments, err := parser.Segments()
	assert.NoError(t, err)
	assert.Len(t, segments, 37)
}

func TestParser_Segments(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		file, err := os.Open("./test/850_with_tilde_segment_terminator.txt")
		assert.NoError(t, err)
		parser := NewParser(file)
		segments, err := parser.Segments()
		assert.NoError(t, err)
		assert.Len(t, segments, 37)
	})
	t.Run("Success from valid ISA length", func(t *testing.T) {
		reader := strings.NewReader("ISA*00*          *00*          *ZZ*EMEDNYBAT      *ZZ*ETIN           *030219*1140*^*00501*006097493*0*T*:~")
		parser := NewParser(reader)
		segments, err := parser.Segments()
		assert.NoError(t, err)
		assert.Len(t, segments, 1)
		assert.Equal(t, "ISA", segments[0].ID)
	})

	t.Run("Error from ISA too short", func(t *testing.T) {
		reader := strings.NewReader("")
		parser := NewParser(reader)
		_, err := parser.Segments()
		assert.Error(t, err)
	})
}
