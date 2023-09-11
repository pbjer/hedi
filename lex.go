package hedi

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

var (
	ErrInvalidISALength = errors.New("invalid ISA length")
)

type Lexer struct {
	reader io.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		reader: reader,
	}
}

// Tokens returns a slice of tokens from the reader. It expects
// input of at least the length of a valid ISA segment (106 bytes),
// or it will error. Accuracy of the returned token types is
// dependent on the ISA.
func (l *Lexer) Tokens() ([]Token, error) {
	isaTokens, separators, err := lexISA(l.reader)
	if err != nil {
		return []Token{}, err
	}
	return append(isaTokens, lexSegments(l.reader, separators)...), nil
}

// lexISA returns a slice of tokens and the defined separators
// for the given input. It expects input of at least 106 bytes,
// or it will error. It does not otherwise validate the ISA or
// the returned tokens.
func lexISA(reader io.Reader) ([]Token, Delimiters, error) {
	isaBuffer := make([]byte, 106)
	n, err := reader.Read(isaBuffer)
	if err != nil {
		return []Token{}, Delimiters{}, err
	}
	if n != 106 {
		return []Token{}, Delimiters{}, ErrInvalidISALength
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

	var tokens []Token

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

// lexSegments returns a slice of tokens based on the supplied
// separators.
func lexSegments(reader io.Reader, separators Delimiters) []Token {
	var tokens []Token

	scanner := bufio.NewScanner(reader)
	scanner.Split(splitter(separators.Segment))

	for scanner.Scan() {
		segment := scanner.Text()
		tokens = append(tokens, lexSegment(strings.NewReader(segment), separators)...)
	}

	return tokens
}

func lexSegment(reader io.Reader, separators Delimiters) []Token {
	var tokens []Token

	scanner := bufio.NewScanner(reader)
	scanner.Split(splitter(separators.Element))

	scanner.Scan() // First scan should always be the segment identifier
	tokens = append(tokens, Token{Type: SegmentIdentifier, Value: scanner.Text()})

	for scanner.Scan() {
		element := scanner.Text()
		tokens = append(tokens, lexElement(strings.NewReader(element), separators)...)
	}

	return append(tokens, Token{Type: SegmentTerminator, Value: string(separators.Segment)})
}

func lexElement(reader io.Reader, separators Delimiters) []Token {
	var tokens []Token

	scanner := bufio.NewScanner(reader)
	scanner.Split(splitter(separators.SubElement))

	scanner.Scan() // First scan should always be the element
	tokens = append(tokens, Token{Type: ElementDelimiter, Value: string(separators.Element)})

	tokens = append(tokens, Token{Type: ElementValue, Value: scanner.Text()})

	var subElementTokens []Token
	for scanner.Scan() { // Any subsequent scans are sub elements
		subElementTokens = append(subElementTokens, Token{Type: SubElementDelimiter, Value: string(separators.SubElement)})
		subElementTokens = append(subElementTokens, Token{Type: SubElementValue, Value: scanner.Text()})
	}

	return append(tokens, subElementTokens...)
}
