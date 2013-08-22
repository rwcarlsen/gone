package endf

import ()

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

func Parse(ch chan Token, tape *Tape) error {
	// get tape header info
	comment := <-ch
	tapeid := <-ch
	id, err := strconv.Atoi(tapeid.Val)
	if err != nil {
		return fmt.Errorf("endf: invalid tape id %v", tapeid.Val)
	}

	if tape.Id != 0 {
		if tape.Comment != comment.Val || tape.Id != id {
			return errors.New("endf: incompatible tapes")
		}
	}

	tape.Id = id
	tape.Comment = comment

	// parse materials
	for done, err := parseMat(ch, tape) {
		if err != nil {
			return err
		} else if done {
			break
		}
	}

	// check for closing record
	_ := <-ch
	mat := <-ch
	id, err := strconv.Atoi(mat.Val)
	if err != nil {
		return errors.New("endf: invalid tape end")
	} else if id != -1 {
		return fmt.Errorf("endf: expected MAT id of -1, got %v", id)
	}
	mf := <-ch
	mt := <-ch
	eof := <-ch
	if eof.Type != TokEOF {
		return fmt.Errorf("endf: expected EOF, found %v", eof.Type)
	}
	return nil
}

func parseMat(ch chan Token tape *Tape) (done bool, err error) {
	// get material header info
	_ = <-ch
	mat := <-ch
	id, err := strconv.Atoi(mat.Val)
	if err != nil {
		return fmt.Errorf("endf: invalid MAT id %v", mat.Val)
	}

	if id == -1 {
		return true, nil
	}

	tape.Id = id
	tape.Comment = comment
	for {
		matid := parseSection(ch, tape)
	}

	// check for closing record
	_ := <-ch
	mat := <-ch
	id, err := strconv.Atoi(mat.Val)
	if err != nil {
		return errors.New("endf: invalid material end")
	} else if id != 0 {
		return fmt.Errorf("endf: expected MAT id of 0, got %v", id)
	}
	mf := <-ch
	mt := <-ch
}
