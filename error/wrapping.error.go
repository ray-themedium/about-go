package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// %w 서식문자를 통해서 에러를 감싸고 새로운 에러를 정의
func MultipleFromString(str string) (int, error) {
	// string 타입을 io.Reader 타입으로 변환 후 bufio.Scanner 생성
	scan := bufio.NewScanner(strings.NewReader(str))
	// 한 단어씩 끊어서 읽을 수 있도록 선언
	scan.Split(bufio.ScanWords)

	pos := 0
	a, n, err := readNextInt(scan)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos: %d err: %w", pos, err)
	}
	pos += n + 1
	b, n, err := readNextInt(scan)
	if err != nil {
		return 0, fmt.Errorf("Failed to readNextInt(), pos: %d err: %w", pos, err)
	}
	return a * b, nil
}

func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	// Split으로 지정한 요청에 맞게 읽어옴
	if !scanner.Scan() {
		return 0, 0, errors.New("Failed to scan")
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word: %s err: %w", word, err)
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		// 감싸진 에러를 다시 꺼내서 변환
		// %w 서식 문자로 감싸진 에러를 꺼내서 변환할 수 있음
		if errors.As(err, &numError) {
			fmt.Println("NumberError:", numError)
		}
	}
}

func main() {
	readEq("123 3")
	readEq("123 abc")
}
