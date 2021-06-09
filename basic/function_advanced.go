package main

import "fmt"

func add(a, b int) int { return a+b } 
func mul(a, b int) int { return a*b }

func getOperator(op string) func (int, int) int {
	if op == "+" {
		return add 
	} else if op == "*" {
		return mul 
	} else {
		return nil
	}
}

func main() {
	var op func (int, int) int
	op = getOperator("*")

	fmt.Println(op(3, 5))
}



// 파일 생성 후 출력, defer 사용
// func main() {
// 	f, err := os.Create("test.txt")
// 	if err != nil {
// 		fmt.Println("Failed to create a file")
// 		return
// 	}

// 	defer fmt.Println("반드시 호출됩니다.")
// 	defer f.Close()
// 	defer fmt.Println("파일을 닫았습니다.") // defer는 역순으로 호출됨

// 	fmt.Println("파일에서 ...을 씁니다.")
// 	fmt.Fprintln(f, "testtest")
// }