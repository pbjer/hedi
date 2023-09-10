package hedi

import (
	"strings"
)

type Element struct {
	Value       string
	SubElements []string
}

func (e *Element) String() string {
	return e.DString(DefaultDelimiters)
}

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

func (e *Element) AddSubElement(value string) {
	if e.SubElements == nil {
		e.SubElements = []string{}
	}
	e.SubElements = append(e.SubElements, value)
}

type Elements []Element

func (ee *Elements) Last() *Element {
	if len(*ee) == 0 {
		return &(*ee)[0]
	}
	return &(*ee)[len(*ee)-1]
}
