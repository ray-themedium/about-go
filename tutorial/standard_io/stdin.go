package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin    // command 라인을 읽음
	defer f.Close() // 표준 읽기를 멈추려면 컨트롤 D

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}
