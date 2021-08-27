package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE"))
	f("%s\n", s.Title("tHis wiLL be A title!"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAlis")) // 대문자든 소문자들 둘이 접었을 때 같은지 확인
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHali"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi")) // 첫 시작(선행자)부터 검사
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is")) // 마지막부터 검사
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "IS"))
	f("Index: %v\n", s.Index("Mihalis", "ha")) // 문자열이 있다면 그 시작점의 index, 없으면 -1
	f("Index: %v\n", s.Index("Mihalis", "Ha"))
	f("Count: %v\n", s.Count("Mihalis", "ha")) // 문자열 개수를 중복없이 리턴
	f("Count: %v\n", s.Count("Mihalisnnnn", "nn"))
	f("Repeat: %v\n", s.Repeat("ab", 6)) // 문자열을 반복하여 새로운 문자열 생성

	f("TrimSpace: %s\n", s.TrimSpace("This\t is a line.\n"))
	f("TrimLeft: %s\n", s.TrimLeft("\tThis is a\t line.\n", "\n\t"))
	f("TrimRight: %s\n", s.TrimRight("\tThis is a\t line.\n", "\n\t"))

	f("Compare: %v\n", s.Compare("Mihalis", "MIHALIS")) // 같으면 0, 아니면 -1, +1
	f("Compare: %v\n", s.Compare("Mihalis", "Mihalis"))
	f("Compare: %v\n", s.Compare("MIHALIS", "MIHalis"))

	f("Fields: %v\n", s.Fields("This is a string!")) // 구분자(space, \n, \t)로 잘라서 slice로 만들어줌
	f("Fields: %v\n", s.Fields("Thisis\na\tstring!"))

	f("%v\n", s.Split("abcd efg", "")) // 직접 구분자를 받아서 slice로 만들어줌

	f("%s\n", s.Replace("abcd efg", "", "_", 2)) // 문자열에서 특정 문자열을 새로운 문자열로 변경하고 int값을 받아 몇 개를 교체할지 선택할 수 있음, -1은 제한 없음

	args := []string{"aa", "bb", "cc"}
	f("Join: %s\n", s.Join(args, "___"))

	trimFunc := func(c rune) bool {
		fmt.Printf("%c\n", c)
		fmt.Println(unicode.IsLetter(c))
		return !unicode.IsLetter(c)
	}

	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t.", trimFunc))
}
