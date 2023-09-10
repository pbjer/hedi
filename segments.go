package hedi

import "strings"

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

func (s *Segments) Last() *Segment {
	if len(*s) == 0 {
		return &(*s)[0]
	}
	return &(*s)[len(*s)-1]
}

func (s *Segments) Add(segments ...Segment) {
	*s = append(*s, segments...)
}
