package hedi

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Segments []Segment

func (s *Segments) String() string {
	return s.DString(DefaultDelimiters)
}

// DString returns a delimited string of the segments.
// Specific delimiters can be optionally defined, where the first
// provided delimiter will be used as a segment terminator, the
// second as the element delimiter, and the third as the sub
// element delimiter. Additional provided delimiters are ignored.
// Delimiters default to '~' (segment), '*' (element), '>' (sub element).
//
// Example usage:
//
//	 segments := Segments{{ ID: "ST", Elements: Elements{{ Value: "850" }, { Value: "0001" }}
//		fmt.Print(segments.DString('|'))
//		Output: ST|850|0001~
func (s *Segments) DString(delimiters Delimiters) string {
	var sb strings.Builder
	for _, segment := range *s {
		sb.WriteString(segment.DString(delimiters))
	}
	return sb.String()
}

func (s *Segments) WriteTo(w io.Writer) (int64, error) {
	return s.DWriteTo(DefaultDelimiters, w)
}

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

func (s *Segments) Last() *Segment {
	if len(*s) == 0 {
		return &(*s)[0]
	}
	return &(*s)[len(*s)-1]
}

func (s *Segments) Add(segments ...Segment) {
	*s = append(*s, segments...)
}
