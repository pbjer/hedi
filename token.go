package hedi

type TokenType string

const (
	SegmentIdentifier   TokenType = "segment_identifier"
	SegmentTerminator   TokenType = "segment_terminator"
	ElementValue        TokenType = "element_value"
	ElementDelimiter    TokenType = "element_delimiter"
	SubElementValue     TokenType = "sub_element_value"
	SubElementDelimiter TokenType = "sub_element_delimiter"
)

type Token struct {
	Type  TokenType `json:"type"`
	Value string    `json:"value"`
}
