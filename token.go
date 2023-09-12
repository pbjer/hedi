package hedi

// TokenType represents the type of token in the parser.
type TokenType string

// Enumerated TokenTypes for various elements in the segment structure.
const (
	// SegmentIdentifier represents the type of token that identifies a segment.
	SegmentIdentifier TokenType = "segment_identifier"

	// SegmentTerminator represents the type of token that terminates a segment.
	SegmentTerminator TokenType = "segment_terminator"

	// ElementValue represents the type of token that holds the value of an element.
	ElementValue TokenType = "element_value"

	// ElementDelimiter represents the type of token that delimits elements.
	ElementDelimiter TokenType = "element_delimiter"

	// SubElementValue represents the type of token that holds the value of a sub-element.
	SubElementValue TokenType = "sub_element_value"

	// SubElementDelimiter represents the type of token that delimits sub-elements.
	SubElementDelimiter TokenType = "sub_element_delimiter"
)

// Token structure holding the type and value of a parsed token.
type Token struct {
	Type  TokenType `json:"type"`
	Value string    `json:"value"`
}
