package hedi

var DefaultDelimiters = Delimiters{
	Segment:    '~',
	Element:    '*',
	SubElement: '~',
}

type Delimiters struct {
	Segment    rune
	Element    rune
	SubElement rune
}

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
