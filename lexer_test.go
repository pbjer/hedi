package hedi

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
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

func TestLexer_Tokens(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		file, err := os.Open("./test/850_with_tilde_segment_terminator.txt")
		assert.NoError(t, err)
		defer file.Close()
		lexer := NewLexer(file)
		tokens, err := lexer.Tokens()
		assert.NoError(t, err)
		assert.Len(t, tokens, 503)
	})
	tests := []struct {
		name    string
		input   string
		wantErr error
		len     int
	}{
		{
			name:    "ValidInput",
			input:   "ISA*00*          *00*          *ZZ*SENDER         *ZZ*RECEIVER       *190430*1230*U*00401*000000000*0*T*|\n",
			wantErr: nil,
			len:     33,
		},
		{
			name:    "InvalidISALength",
			input:   "ISA*00*          *00*",
			wantErr: ErrInvalidISALength,
			len:     0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lexer := NewLexer(strings.NewReader(tt.input))
			tokens, err := lexer.Tokens()

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Len(t, tokens, tt.len)
		})
	}
}
