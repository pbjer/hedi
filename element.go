package hedi

import (
	"strings"
)

// Element represents an individual EDI element, containing a value and optional sub-elements.
type Element struct {
	Value       string
	SubElements []string
}

// String returns the default delimited string representation of the Element.
// It uses the DefaultDelimiters for formatting.
func (e *Element) String() string {
	return e.DString(DefaultDelimiters)
}

// DString returns a delimited string representation of the Element.
// It formats the Element's value and sub-elements using the provided Delimiters.
func (e *Element) DString(delimiters Delimiters) string {
	var sb strings.Builder

	sb.WriteRune(delimiters.Element)
	sb.WriteString(e.Value)
	for _, subElement := range e.SubElements {
		sb.WriteRune(delimiters.SubElement)
		sb.WriteString(subElement)
	}

	return sb.String()
}

// AddSubElement appends a sub-element value to the Element's SubElements slice.
// Initializes SubElements if it is nil.
func (e *Element) AddSubElement(value string) {
	if e.SubElements == nil {
		e.SubElements = []string{}
	}
	e.SubElements = append(e.SubElements, value)
}

// Elements is a slice of Element structs, often representing a list of elements in an EDI segment.
type Elements []Element

// Last returns the last Element in the Elements slice.
// Returns nil and false if the Elements slice is empty.
func (ee *Elements) Last() (*Element, bool) {
	if len(*ee) == 0 {
		return nil, false
	}
	return &(*ee)[len(*ee)-1], true
}
