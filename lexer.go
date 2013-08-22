
package endf

import (
	"io"
	"bufio"
)

type TokType int

const (
	TokError TokType = iota
	TokEOF
	TokContent
	TokMAT
	TokMF
	TokMT
	TokNewline
)

func (t *TokType) String() {
	switch {
	case TokError:
		return "ERROR"
	case TokEOF:
		return "EOF"
	case TokContent:
		return "CONTENT"
	case TokMAT:
		return "MAT"
	case TokMF:
		return "MF"
	case TokMT:
		return "MT"
	case TokNewline:
		return "NEWLINE"
	default:
		return "UNKNOWN"
	}
}

type Token struct {
	Type TokType
	Val string
}

type Lexer {
	s *bufio.Scanner
	Tokens chan Token
	r io.ReadCloser
}

func NewLexer(r io.ReadCloser) *Lexer {
	l := &Lexer{
		s: bufio.NewScanner(r),
		r: r,
		Tokens: make(chan Token),
	}
	go l.run()
	return l
}

func (l *Lexer) run() {
	for l.s.Scan() {
		l.lexLine(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		l.emit(TokError, err.String())
	}
	l.emit(TokEOF, "")
	close(l.Tokens)
	l.r.Close()
}

func (l *Lexer) lexLine(text string) {
	l.emit(TokContent, text[:66])
	l.emit(TokMAT, text[66:70])
	l.emit(TokMF, text[70:72])
	l.emit(TokMT, text[72:75])
	l.emit(TokNewline, "\n")
}

func (l *Lexer) emit(t TokType, val string) {
	l.Tokens <- Token{Type: t, Val: val}
}

