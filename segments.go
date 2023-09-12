package hedi

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Segments is a slice of Segment types.
type Segments []Segment

// String satisfies the fmt.Stringer interface, delegating to DString.
func (s *Segments) String() string {
	return s.DString(DefaultDelimiters)
}

// DString constructs a string representation of Segments using provided delimiters.
func (s *Segments) DString(delimiters Delimiters) string {
	var sb strings.Builder
	for _, segment := range *s {
		sb.WriteString(segment.DString(delimiters))
	}
	return sb.String()
}

// WriteTo satisfies the io.WriterTo interface, delegating to DWriteTo.
func (s *Segments) WriteTo(w io.Writer) (int64, error) {
	return s.DWriteTo(DefaultDelimiters, w)
}

// DWriteTo writes the Segments to an io.Writer w, formatted with specified delimiters.
// Returns the number of bytes written and any error encountered.
func (s *Segments) DWriteTo(d Delimiters, w io.Writer) (int64, error) {
	var total int64
	bufferedWriter := bufio.NewWriter(w)

	for _, segment := range *s {
		n, err := bufferedWriter.WriteString(segment.ID)
		total += int64(n)
		if err != nil {
			return total, err
		}

		for _, element := range segment.Elements {
			m, err := bufferedWriter.WriteString(fmt.Sprintf("%c%s", d.Element, element.Value))
			total += int64(m)
			if err != nil {
				return total, err
			}

			for _, sub := range element.SubElements {
				o, err := bufferedWriter.WriteString(fmt.Sprintf("%c%s", d.SubElement, sub))
				total += int64(o)
				if err != nil {
					return total, err
				}
			}
		}

		p, err := bufferedWriter.WriteString(string(d.Segment))
		total += int64(p)
		if err != nil {
			return total, err
		}

	}

	if err := bufferedWriter.Flush(); err != nil {
		return total, err
	}

	return total, nil
}

// Last returns a pointer to the last segment in the list, or nil if the list is empty.
// The boolean return value indicates the presence of a last segment.
func (s *Segments) Last() (*Segment, bool) {
	if len(*s) == 0 {
		return nil, false
	}
	return &(*s)[len(*s)-1], true
}
