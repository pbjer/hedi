# Hedi

Hedi is a library for reading and writing Electronic Data Interchange (EDI) messages.

## Installation
```bash
go get github.com/pbjer/hedi
```

## Usage
### Parsing
```go
msg := "ISA*01*0000000000*01*0000000000*ZZ*ABCDEFGHIJKLMNO*ZZ*123456789012345*101127*1719*U*00400*000003438*0*P*>~"
reader := strings.NewReader(msg)
parser := hedi.NewParser(reader)
segments, err := parser.Parse()
if err != nil {
  // ...
}

for _, segment := range segments {
  // ...
}
```

### Lexing
```go
msg := "ISA*01*0000000000*01*0000000000*ZZ*ABCDEFGHIJKLMNO*ZZ*123456789012345*101127*1719*U*00400*000003438*0*P*>~"
reader := strings.NewReader(msg)
lexer := hedi.NewLexer(reader)
tokens, err := lexer.Lex()
if err != nil {
  // ...
}

for _, token := range tokens {
  // ...
}
```

### Marshaling
```go
segments := hedi.Segments{{
  ID: "ST",
  Elements: hedi.Elements{{ Value: "850" }, { Value: "000000010" }},
}}

fmt.Println(segments.String())
// ST*850*000000010~
```