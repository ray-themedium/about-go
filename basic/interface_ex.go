package main

type ReadWriter interface {
	Read()
	Write()
}

type File struct {}
func (f *File) Read() {}
func (f *File) Write() {}

func ReadWrite(rw ReadWriter) {
	rw.Read()
	rw.Write()
}

// ----------
type Stringer interface {
	String()
}

type Reader interface {
	Read()
}

func CheckAndRun(stringer Stringer) {
	if r, ok := stringer.(Reader); ok {
		r.Read()
	}
}

func main() {
	f := &File{}
	ReadWrite(f)
}