package hedi

// DefaultDelimiters defines the default Delimiters used in EDI files.
var DefaultDelimiters = Delimiters{
	Segment:    '~',
	Element:    '*',
	SubElement: '~',
}

// Delimiters contains the delimiters used for splitting segments, elements,
// and sub-elements in EDI files.
type Delimiters struct {
	Segment    rune
	Element    rune
	SubElement rune
}

// splitter returns a bufio.SplitFunc function for use in bufio.Scanner.
// The returned function splits the data into tokens separated by the given rune.
// This is commonly used for parsing segments, elements, or sub-elements in EDI files.
func splitter(separator rune) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == byte(separator) {
				return i + 1, data[:i], nil
			}
		}
		if atEOF && len(data) != 0 {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
}
