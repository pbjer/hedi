package hedi

import (
	"strings"
)

// Segment represents an EDI segment, which consists of an ID and a list of Elements.
type Segment struct {
	ID       string
	Elements Elements
}

// NewSegment constructs a new Segment with the given ID.
func NewSegment(id string) *Segment {
	return &Segment{ID: id}
}

// String converts the Segment to its EDI string representation using default delimiters.
func (s *Segment) String() string {
	return s.DString(DefaultDelimiters)
}

// DString converts the Segment to its EDI string representation using the provided delimiters.
func (s *Segment) DString(delimiters Delimiters) string {
	var sb strings.Builder

	// Append Segment ID
	sb.WriteString(s.ID)

	// Append each Element's string representation
	for _, element := range s.Elements {
		sb.WriteString(element.DString(delimiters))
	}

	// Append Segment delimiter
	sb.WriteRune(delimiters.Segment)

	return sb.String()
}

// GetElement retrieves the Element at the specified index within the Segment.
// Returns the Element and a boolean indicating whether the Element was found.
func (s *Segment) GetElement(index int) (Element, bool) {
	if len(s.Elements) <= index {
		return Element{}, false
	}
	return s.Elements[index], true
}

// AddElement appends an Element to the end of the Segment.
func (s *Segment) AddElement(element Element) {
	s.SetElement(len(s.Elements), element)
}

// SetElement replaces or appends an Element at the specified index in the Segment.
// If the index exceeds the current size, the Elements slice is expanded.
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
