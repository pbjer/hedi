package hedi

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

var (
	// ErrInvalidISALength represents an error for invalid ISA segment length.
	ErrInvalidISALength = errors.New("invalid ISA length")
)

// Lexer wraps an io.Reader for lexing EDI files.
type Lexer struct {
	reader io.Reader
}

// NewLexer initializes a new Lexer with a given io.Reader.
func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		reader: reader,
	}
}

// Tokens lexes the input and returns a slice of Token structs.
// It expects an input that starts with a valid ISA segment of 106 bytes.
// Returns an error if the input does not meet the criteria.
func (l *Lexer) Tokens() ([]Token, error) {
	isaTokens, separators, err := lexISA(l.reader)
	if err != nil {
		return []Token{}, err
	}
	return append(isaTokens, lexSegments(l.reader, separators)...), nil
}

// lexISA tokenizes the ISA segment and returns the identified delimiters.
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

	// Record segment identifier and consumed delimiter token
	tokens = append(tokens,
		Token{Type: SegmentIdentifier, Value: segmentParts[0]},
		Token{Type: ElementDelimiter, Value: string(separators.Element)},
	)

	// Record each element in the segment, starting after the identifier, ending before the sub element delimiter value
	for i, part := range segmentParts[1 : len(segmentParts)-1] {
		tokens = append(tokens, Token{Type: ElementValue, Value: part})
		if i < len(segmentParts)-2 { // Add consumed delimiter if not the last element
			tokens = append(tokens, Token{Type: ElementDelimiter, Value: string(separators.Element)})
		}
	}

	// Record the sub element delimiter value and segment terminator
	tokens = append(tokens,
		Token{Type: ElementValue, Value: string(separators.SubElement)},
		Token{Type: SegmentTerminator, Value: string(separators.Segment)},
	)

	return tokens, *separators, nil
}

// lexSegments tokenizes all subsequent segments in the EDI file using the delimiters identified in the ISA segment.
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

// lexSegment tokenizes a single segment using the provided delimiters.
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

// lexElement tokenizes an element and its sub-elements, if any, using the provided delimiters.
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

// splitter returns a bufio.SplitFunc function for use in bufio.Scanner.
// The returned function splits the data into tokens separated by the given rune.
// This is commonly used for parsing segments, elements, or sub-elements in EDI files.
func splitter(separator rune) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == byte(separator) {
				return i + 1, data[:i], nil
			}
		}
		if atEOF && len(data) != 0 {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
}
