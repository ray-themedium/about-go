package main

import (
	"about-golang/testpkg/publicpkg"
	"fmt"
)

func main() {
	fmt.Println("PI : ", publicpkg.PI)
	publicpkg.PublicFunc()
	
	var myint publicpkg.MyInt = 10
	fmt.Println("myint: ", myint)

	var mystruct = publicpkg.MyStruct{Age: 18}
	fmt.Println("mystruct:", mystruct)
}