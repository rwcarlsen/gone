package endf

import (
	"errors"
)

type Tape struct {
	Id        int
	Comment   string
	Materials []*Material
}

type Material struct {
	Id    int // MAT
	Files []*File
}

type File struct {
	Id       int // MF
	Sections []*Section
}

type Section struct {
	Id      int // MT
	Records []string
}

func Parse(ch chan *Record) (*Tape, error) {
	header, ok := <-ch
	if !ok {
		return nil, errors.New("endf: Unexpected EOF")
	}

	mats, err := buildMats(ch)
	if err != nil {
		return nil, err
	}

	tape := &Tape{Id: header.MAT, Comment: header.Content, Materials: mats}
	return tape, nil
}

func buildMats(ch chan *Record) ([]*Material, error) {
	currMat := new(Material)
	currFile := new(File)
	currSec := new(Section)
	mats := make([]*Material, 0)
	for r := range ch {
		if r.MAT == -1 {
			return mats, nil
		} else if r.MAT == 0 {
			mats = append(mats, currMat)
			currMat = new(Material)
		} else if r.MF == 0 {
			currMat.Files = append(currMat.Files, currFile)
			currFile = new(File)
		} else if r.MT == 0 {
			currFile.Sections = append(currFile.Sections, currSec)
			currSec = new(Section)
		} else {
			currSec.Id = r.MT
			currSec.Records = append(currSec.Records, r.Content)
		}
	}

	return nil, errors.New("endf: Unexpected EOF")
}
