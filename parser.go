package endf

import (
	"strconv"
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

// intermediate type for parsing
type record struct {
	Content string
	MAT     int
	MF      int
	MT      int
}

func Parse(ch chan Token) (*Tape, error) {
	// construct record entries
	records := make([]*record, 0)
	for {
		text := <-ch
		if text.Type == TokEOF {
			break
		}

		mat := <-ch
		if mat.Type == TokEOF {
			return nil, errors.New("endf: unexpected EOF")
		}
		matid, err := strconv.Atoi(mat)
		if err != nil {
			return nil, fmt.Errorf("endf: %v is not a valid id")
		}

		mf := <-ch
		if mat.Type == TokEOF {
			return nil, errors.New("endf: unexpected EOF")
		}
		mfid, err := strconv.Atoi(mf)
		if err != nil {
			return nil, fmt.Errorf("endf: %v is not a valid id")
		}

		mt := <-ch
		if mat.Type == TokEOF {
			return nil, errors.New("endf: unexpected EOF")
		}
		mtid, err := strconv.Atoi(mt)
		if err != nil {
			return nil, fmt.Errorf("endf: %v is not a valid id")
		}

		records = append(records, &record{text, matid, mfid, mtid})
	}

	header := records[0]
	mats := buildMats(records[1:])
	tape := &Tape{Id: header.MAT, Comment: header.Content, Materials: mats}
	return tape, nil
}

func buildMats(records []*record) []*Material {
	currMat := make(Material)
	currFile := make(File)
	currSec := make(Section)
	mats := make([]*Material, 0)
	for _, r := range records {
		if r.MAT == -1 {
			break // end of tape
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

	return mats
}
