package main

import "fmt"

type Stringer interface {
	String() string
}

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type A struct{ a int }

func (a *A) String() string {
	return fmt.Sprintf("%d", a.a)
}

type File struct{}

func (f *File) Read()  {}
func (f *File) Close() {}

// Stringer 인터페이스 타입을 구체화된 타입인 *A 타입으로 변환
func PrintA(s Stringer) {
	// 변환한 결과, 변환 성공 여부
	b, ok := s.(*A)
	if !ok {
		fmt.Println("failed")
	} else {
		fmt.Println(b.a)
	}
}

// Reader 인터페이스 타입을 Closer 인터페이스 타입으로 변환
func ReadFile(reader Reader) {
	if c, ok := reader.(Closer); ok {
		c.Close()
	}
}

func main() {
	a := &A{100}
	PrintA(a)

	file := &File{}
	ReadFile(file)
}
