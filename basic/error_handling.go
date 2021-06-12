package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line)
	return err
}

const filename string = "data.txt"

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("제곱근은 양수여야 합니다. f:%g", f)
	}
	return math.Sqrt(f), nil
}

type PasswordError struct {
	Len int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{ len(password), 8 }
	}
	return nil
}

func main() {
	// 1
	// line, err := ReadFile(filename)
	// if err != nil {
	// 	err = WriteFile(filename, "This is WriteFile")
	// 	if err != nil {
	// 		fmt.Println("파일 생성에 실패했습니다.", err)
	// 		return
	// 	}
	// 	line, err = ReadFile(filename)
	// 	if err != nil {
	// 		fmt.Println("파일 읽기에 실패했습니다.", err)
	// 		return
	// 	}
	// }
	// fmt.Println("파일내용 :", line)

	// 2
	// sqrt, err := Sqrt(-2)
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// }
	// fmt.Println(sqrt)

	// 3
	err := RegisterAccount("myId", "myPw")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d RequireLen:%d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	}else {
		fmt.Println("회원 가입됐습니다.")
	}
}