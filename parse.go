package hedi

import (
	"errors"
	"io"
)

var (
	ErrSegmentIdentifierExpected = errors.New("segment identifier expected")
	ErrElementExpected           = errors.New("element expected")
)

type Parser struct {
	reader io.Reader
}

func NewParser(reader io.Reader) *Parser {
	return &Parser{
		reader: reader,
	}
}

// Segments parses Segments from the underlying reader.
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
