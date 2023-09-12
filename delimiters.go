package hedi

// DefaultDelimiters defines the default Delimiters used in EDI files.
var DefaultDelimiters = Delimiters{
	Segment:    '~',
	Element:    '*',
	SubElement: '>',
}

// Delimiters contains the delimiters used for splitting segments, elements,
// and sub-elements in EDI files.
type Delimiters struct {
	Segment    rune
	Element    rune
	SubElement rune
}
