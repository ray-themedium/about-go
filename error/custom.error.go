package main

import (
	"errors"
	"fmt"
)

// 에러 타입 선언
type MyError struct {
	Len        int
	RequireLen int
}

// error 인터페이스 구현
func (myError MyError) Error() string {
	return "myError"
}

func validation(data string) error {
	if len([]rune(data)) < 10 {
		return MyError{len([]rune(data)), 10}
	}
	return nil
}

func ExampleCustomError1() error {
	return fmt.Errorf("error1")
}

func ExampleCustomError2() error {
	return errors.New("error2")
}

func main() {
	myError := validation("ㅋㅋㅋㅋㅋㅋ")
	if myError != nil {
		if errInfo, ok := myError.(MyError); ok {
			fmt.Printf("%v len:%d required:%d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	}

	err1 := ExampleCustomError1()
	err2 := ExampleCustomError2()

	if err1 != nil {
		fmt.Println(err1)
	}
	if err2 != nil {
		fmt.Println(err2)
	}
}
