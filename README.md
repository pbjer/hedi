# Hedi

Hedi is a library for interacting with Electronic Data Interchange (EDI) messages.

[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/pbjer/hedi)](https://pkg.go.dev/mod/github.com/pbjer/hedi)
![Tests](https://github.com/pbjer/hedi/actions/workflows/main.yml/badge.svg?branch=main&kill_cache=1)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
___

## Installation
```bash
go get github.com/pbjer/hedi
```

## Usage
### Lexing
```go
msg := "ISA*01*0000000000*01*0000000000*ZZ*ABCDEFGHIJKLMNO*ZZ*123456789012345*101127*1719*U*00400*000003438*0*P*>~"
reader := strings.NewReader(msg)
lexer := hedi.NewLexer(reader)
tokens, err := lexer.Tokens()
if err != nil {
  // ...
}
```
### Parsing
```go
msg := "ISA*01*0000000000*01*0000000000*ZZ*ABCDEFGHIJKLMNO*ZZ*123456789012345*101127*1719*U*00400*000003438*0*P*>~"
reader := strings.NewReader(msg)
parser := hedi.NewParser(reader)
segments, err := parser.Segments()
if err != nil {
  // ...
}
```

### Serialization

#### Stringer
Hedi's EDI types implement the `String() string` stringer interface for simple string serialization.

To override the default delimiters, you can use `DString(d hedi.Delimiters) string`, which Stringer depends on under the hood.
```go
segments := hedi.Segments{{
  ID: "ST",
  Elements: hedi.Elements{{ Value: "850" }, { Value: "000000010" }},
}}

fmt.Println(segments)
// ST*850*000000010~

delimiters := hedi.Delimeters{
  Segment: '\n',
  Element: '|',
  SubElement: '>',
}
	
fmt.Println(segments.DString(delimiters))
// ST|850|000000010
//
```

#### WriterTo
Hedi's `Segments` EDI type implements the WriterTo interface for efficient string serialization to an `io.Writer`.

To override the default delimiters, you can use `DWriteTo(d hedi.Delimiters, w io.Writer) (int64, error)`, which WriteTo depends on under the hood.
```go
file, _ := os.Create("850.txt")
if err != nil {
  // ...
}
defer file.Close()

segments := hedi.Segments{{
  ID: "ST",
  Elements: hedi.Elements{{ Value: "850" }, { Value: "000000010" }},
}}

_, err = segments.WriteTo(file)
if err != nil {
  // ...
}
```
