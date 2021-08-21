// 커맨드라인 인수 중에서 올바른 숫자 더하기
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("인수가 없음")
		os.Exit(1)
	}

	total := 0.0
	numberCount := 0

	for i := 1; i < len(args); i++ {
		n, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			continue
		}
		numberCount++
		total += n
	}

	fmt.Println("총합:", total)
	fmt.Println("평균:", total/float64(numberCount))
}
