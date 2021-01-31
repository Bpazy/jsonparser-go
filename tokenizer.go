package jsonparser

import (
	"strings"
)

type TokenType int

const (
	BeginObject TokenType = iota
	EndObject
	BeginArray
	EndArray
	NULL
	Number
	String
	Bool
	SepColon
	SepComma
	EndDocument
)

type Token struct {
	TokenType TokenType
	Value     string
}

func newToken(tokenType TokenType, value string) Token {
	return Token{TokenType: tokenType, Value: value}
}

type Tokenizer struct {
	Tokens []Token
	reader *strings.Reader
}

func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer{
		Tokens: []Token{},
		reader: strings.NewReader(input),
	}
}

// Tokenize tokenize json string
func (t *Tokenizer) Tokenize() {
	for {
		token := t.nextToken()
		if token.TokenType == EndDocument {
			break
		}
		t.Tokens = append(t.Tokens, token)
	}
}

func (t *Tokenizer) nextToken() Token {
	b, err := t.nextNoneSpaceChar()
	if err != nil {
		return newToken(EndDocument, "")
	}

	switch b {
	case '{':
		return newToken(BeginObject, "{")
	case '}':
		return newToken(EndObject, "}")
	case '[':
		return newToken(BeginArray, "[")
	case ']':
		return newToken(EndArray, "]")
	case ',':
		return newToken(SepComma, ",")
	case ':':
		return newToken(SepColon, ":")
	case 'n':
		return t.readNull()
	case 't':
	case 'f':
		return t.readBool()
	case '"':
		return t.readString()
	}
	return t.readNumber()
}

func (t *Tokenizer) nextNoneSpaceChar() (byte, error) {
	for {
		b, err := t.reader.ReadByte()
		if err != nil {
			return 0, err
		}
		// Jump white space
		if b != ' ' {
			return b, nil
		}
	}
}

func (t *Tokenizer) readNumber() Token {
	mustUnreadByte(t)

	numStr := ""
	for {
		b, err := t.reader.ReadByte()
		if err != nil {
			panic(err)
		}

		if b >= 48 && b <= 57 {
			numStr += string(b)
		} else {
			mustUnreadByte(t)
			return newToken(Number, numStr)
		}
	}
}

func (t *Tokenizer) readNull() Token {
	mustUnreadByte(t)

	// Skip 4 byte
	for i := 0; i < 4; i++ {
		if _, err := t.reader.ReadByte(); err != nil {
			panic(err)
		}
	}

	return newToken(NULL, "null")
}

func (t *Tokenizer) readBool() Token {
	mustUnreadByte(t)

	b, err := t.reader.ReadByte()
	if err != nil {
		panic(err)
	}

	// "true" starts with t
	if b == 't' {
		skip(t.reader, 3)
		return newToken(Bool, "true")
	}

	// "false“ Obviously
	skip(t.reader, 4)
	return newToken(Bool, "false")
}

func (t *Tokenizer) readString() Token {
	s := ""
	for {
		b, err := t.reader.ReadByte()
		if err != nil {
			panic(err)
		}
		if b == '"' {
			return newToken(String, s)
		}
		s += string(b)
	}
}

func skip(r *strings.Reader, n int) {
	for i := 0; i < n; i++ {
		_, _ = r.ReadByte()
	}
}

func mustUnreadByte(t *Tokenizer) {
	err := t.reader.UnreadByte()
	if err != nil {
		panic(err)
	}
}
