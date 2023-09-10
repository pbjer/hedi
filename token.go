package hedi

import (
	"encoding/json"
	"strings"
)

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

type Tokens []Token

func (ts *Tokens) String() string {
	sb := strings.Builder{}
	for _, token := range *ts {
		sb.WriteString(token.Value)
	}
	return sb.String()
}

func (ts *Tokens) Pprint() string {
	sb := strings.Builder{}
	for _, token := range *ts {
		if token.Type == SegmentTerminator {
			sb.WriteString("\n")
			continue
		}
		sb.WriteString(token.Value)
	}
	return sb.String()
}

func (ts *Tokens) JSON() []byte {
	result, err := json.Marshal(ts)
	if err != nil {
		return nil
	}
	return result
}
