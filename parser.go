package hedi

import (
	"errors"
	"io"
)

var (
	// ErrSegmentIdentifierExpected is returned when a segment identifier is expected but not found.
	ErrSegmentIdentifierExpected = errors.New("segment identifier expected")
	// ErrElementExpected is returned when an element is expected but not found.
	ErrElementExpected = errors.New("element expected")
)

// Parser encapsulates the parsing logic for EDI files.
type Parser struct {
	reader io.Reader
}

// NewParser creates a new Parser instance with the given io.Reader.
func NewParser(reader io.Reader) *Parser {
	return &Parser{
		reader: reader,
	}
}

// Segments reads from the underlying reader and converts the token stream into Segments.
// It returns an error if the token stream does not conform to the expected structure.
func (p *Parser) Segments() (Segments, error) {
	lexer := NewLexer(p.reader)
	tokens, err := lexer.Tokens()
	if err != nil {
		return nil, err
	}
	segments := Segments{}
	for _, token := range tokens {
		switch token.Type {
		case SegmentIdentifier:
			segments = append(segments, *NewSegment(token.Value))
		case ElementValue:
			lastSegment, ok := segments.Last()
			if !ok {
				return nil, ErrSegmentIdentifierExpected
			}
			lastSegment.AddElement(Element{Value: token.Value})
		case SubElementValue:
			lastSegment, ok := segments.Last()
			if !ok {
				return nil, ErrSegmentIdentifierExpected
			}
			lastElement, ok := lastSegment.Elements.Last()
			if !ok {
				return nil, ErrElementExpected
			}
			lastElement.AddSubElement(token.Value)
		}
	}
	return segments, nil
}
