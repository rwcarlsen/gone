package endf

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// intermediate type for parsing
type Record struct {
	Content string
	MAT     int
	MF      int
	MT      int
}

type Lexer struct {
	s       *bufio.Scanner
	Records chan *Record
	r       io.ReadCloser
}

func NewLexer(r io.ReadCloser) *Lexer {
	l := &Lexer{
		s:       bufio.NewScanner(r),
		r:       r,
		Records: make(chan *Record),
	}
	go l.run()
	return l
}

func (l *Lexer) run() {
	for l.s.Scan() {
		l.lexLine(l.s.Text())
	}
	if err := l.s.Err(); err != nil {
		l.Records <- &Record{Content: fmt.Sprintf("ERROR: %v", err)}
	}
	close(l.Records)
	l.r.Close()
}

func (l *Lexer) lexLine(text string) {
	content := text[:66]
	mat := must(strconv.Atoi(strings.TrimSpace(text[66:70])))
	mf := must(strconv.Atoi(strings.TrimSpace(text[70:72])))
	mt := must(strconv.Atoi(strings.TrimSpace(text[72:75])))

	r := &Record{
		Content: content,
		MAT:     mat,
		MF:      mf,
		MT:      mt,
	}

	l.Records <- r
}

func must(v int, err error) int {
	if err != nil {
		panic(err)
	}
	return v
}
