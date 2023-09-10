package hedi

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

type Lexer struct {
	reader io.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		reader: reader,
	}
}

// Lex returns a slice of tokens for the given input. It expects
// input of at least the length of a valid ISA segment (106 bytes),
// or it will error. Accuracy of the returned token types is
// dependent on the ISA.
func (l *Lexer) Lex() (Tokens, error) {
	isaTokens, separators, err := LexISA(l.reader)
	if err != nil {
		return Tokens{}, err
	}
	return append(isaTokens, LexSegments(l.reader, separators)...), nil
}

// LexISA returns a slice of tokens and the defined separators
// for the given input. It expects input of at least 106 bytes,
// or it will error. It does not otherwise validate the ISA or
// the returned tokens.
func LexISA(reader io.Reader) (Tokens, Delimiters, error) {
	isaBuffer := make([]byte, 106)
	n, err := reader.Read(isaBuffer)
	if err != nil {
		return Tokens{}, Delimiters{}, err
	}
	if n != 106 {
		return Tokens{}, Delimiters{}, errors.New("invalid ISA length")
	}

	isaString := string(isaBuffer)
	elementSeparator := isaString[103]
	subElementSeparator := isaString[104]
	segmentSeparator := isaString[105]

	separators := &Delimiters{
		Segment:    rune(segmentSeparator),
		Element:    rune(elementSeparator),
		SubElement: rune(subElementSeparator),
	}

	var tokens Tokens

	// Split the segment into its identifier and elements
	segmentParts := strings.Split(isaString, string(separators.Element))

	// Record segment identifier and consumed delimiter tokens
	tokens = append(tokens,
		Token{Type: SegmentIdentifier, Value: segmentParts[0]},
		Token{Type: ElementDelimiter, Value: string(separators.Element)},
	)

	// Record each element in the segment, starting after the identifier, ending before the segment terminator
	for i, part := range segmentParts[1 : len(segmentParts)-1] {
		tokens = append(tokens, Token{Type: ElementValue, Value: part})
		if i < len(segmentParts)-2 { // Add consumed delimiter if not the last element
			tokens = append(tokens, Token{Type: ElementDelimiter, Value: string(separators.Element)})
		}
	}

	// Record the segment terminator
	tokens = append(tokens, Token{Type: SegmentTerminator, Value: string(separators.Segment)})

	return tokens, *separators, nil
}

// LexSegments returns a slice of tokens based on the supplied
// separators.
func LexSegments(reader io.Reader, separators Delimiters) Tokens {
	var tokens Tokens

	scanner := bufio.NewScanner(reader)
	scanner.Split(splitter(separators.Segment))

	for scanner.Scan() {
		segment := scanner.Text()
		tokens = append(tokens, LexSegment(strings.NewReader(segment), separators)...)
	}

	return tokens
}

func LexSegment(reader io.Reader, separators Delimiters) Tokens {
	var tokens Tokens

	scanner := bufio.NewScanner(reader)
	scanner.Split(splitter(separators.Element))

	scanner.Scan() // First scan should always be the segment identifier
	tokens = append(tokens, Token{Type: SegmentIdentifier, Value: scanner.Text()})

	for scanner.Scan() {
		element := scanner.Text()
		tokens = append(tokens, LexElement(strings.NewReader(element), separators)...)
	}

	return append(tokens, Token{Type: SegmentTerminator, Value: string(separators.Segment)})
}

func LexElement(reader io.Reader, separators Delimiters) Tokens {
	var tokens Tokens

	scanner := bufio.NewScanner(reader)
	scanner.Split(splitter(separators.SubElement))

	scanner.Scan() // First scan should always be the element
	tokens = append(tokens, Token{Type: ElementDelimiter, Value: string(separators.Element)})

	tokens = append(tokens, Token{Type: ElementValue, Value: scanner.Text()})

	subElementTokens := Tokens{}
	for scanner.Scan() { // Any subsequent scans are sub elements
		subElementTokens = append(subElementTokens, Token{Type: SubElementDelimiter, Value: string(separators.SubElement)})
		subElementTokens = append(subElementTokens, Token{Type: SubElementValue, Value: scanner.Text()})
	}

	return append(tokens, subElementTokens...)
}
