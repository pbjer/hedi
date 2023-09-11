package hedi

import "io"

type Parser struct {
	reader io.Reader
}

func NewParser(reader io.Reader) *Parser {
	return &Parser{
		reader: reader,
	}
}

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
			segments.Add(*NewSegment(token.Value))
		case ElementValue:
			segments.Last().AddElement(Element{Value: token.Value})
		case SubElementValue:
			segments.Last().Elements.Last().AddSubElement(token.Value)
		}
	}
	return segments, nil
}
