package endf

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	f, err := os.Open("neutrons/n-001_H_001.endf")
	if err != nil {
		t.Fatal(err)
	}

	l := NewLexer(f)
	tape, err := Parse(l.Records)
	if err != nil {
		t.Fatal(err)
	}

	mat := tape.Materials[0]
	file := mat.Files[0]
	sec := file.Sections[0]

	t.Logf("material id (MAT): %v", mat.Id)
	t.Logf("file id (MF): %v", file.Id)
	t.Logf("section id (MT): %v", sec.Id)
	for _, r := range sec.Records {
		t.Log(r)
	}
}
