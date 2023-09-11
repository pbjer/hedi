package hedi

import (
	"strings"
)

type Segment struct {
	ID       string
	Elements Elements
}

func NewSegment(id string) *Segment {
	return &Segment{ID: id}
}

// String returns a string of the segment.
func (s *Segment) String() string {
	return s.DString(DefaultDelimiters)
}

// DString returns a string of the segment formatted using the provided delimiters.
func (s *Segment) DString(delimiters Delimiters) string {
	var sb strings.Builder

	sb.WriteString(s.ID)
	for _, element := range s.Elements {
		sb.WriteString(element.DString(delimiters))
	}
	sb.WriteRune(delimiters.Segment)

	return sb.String()
}

// GetElement returns the element at the specified index.
// The second return value will be true if an Element is
// found, otherwise false.
func (s *Segment) GetElement(index int) (Element, bool) {
	if len(s.Elements) <= index {
		return Element{}, false
	}
	return s.Elements[index], true
}

// AddElement adds an element at the end of the current segment.
func (s *Segment) AddElement(element Element) {
	s.SetElement(len(s.Elements), element)
}

// SetElement assigns the provided element at the specified
// index in the Segment. If the index is out of the current
// range, the Elements slice is dynamically expanded to
// accommodate the new element.
func (s *Segment) SetElement(index int, element Element) {
	delta := 0
	if len(s.Elements) <= index {
		delta = index - len(s.Elements)
		if delta == 0 {
			delta = 1
		}
		s.Elements = append(s.Elements, make([]Element, delta)...)
	}
	s.Elements[index] = element
}
