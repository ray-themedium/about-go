package main

import "fmt"

type Reader interface {
	Read() (n int, err error)
	Close() error
}

type Writer interface {
	Write() (n int, err error)
	Close() error
}

// 포함된 인터페이스
// Read(), Wirte(), Close() 메서드를 포함
type ReadWriter interface {
	Reader
	Writer
}

// 빈 인터페이스가 매개변수인 함수
func PrintValue(v interface{}) {
	fmt.Println(v)
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", t)
	case float64:
		fmt.Printf("v is float64 %f\n", t)
	default:
		fmt.Printf("Not supported type: %T:%v\n", t, t)
	}
}

func main() {
	PrintValue(10)
	PrintValue(10.0)
	PrintValue("zzz")
}
