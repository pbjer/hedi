package hedi

import (
	"strings"
)

type Element struct {
	Value       string
	SubElements []string
}

// DString returns a string of the element.
func (e *Element) String() string {
	return e.DString(DefaultDelimiters)
}

// DString returns a string of the element with the provided delimiters.
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

// AddSubElement adds a sub-element at the end of the current segment.
func (e *Element) AddSubElement(value string) {
	if e.SubElements == nil {
		e.SubElements = []string{}
	}
	e.SubElements = append(e.SubElements, value)
}

type Elements []Element

// Last returns the last element in an element list.
func (ee *Elements) Last() (*Element, bool) {
	if len(*ee) == 0 {
		return nil, false
	}
	return &(*ee)[len(*ee)-1], true
}
