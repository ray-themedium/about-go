package main

type Student struct {
	Name string
	Age int
}

type Students []Student
func (s Students) Len() int {return len(s)}
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Students) Swap(i, j int) {s[i], s[j] = s[j], s[i]}

func main() {
	// arr := []int{6,5,8,9,1,0,2}
	// sort.Ints(arr)
	// fmt.Println(arr)

	// students := Students {
	// 	{"화랑", 31}, {"백두산", 52}, {"류", 42}, {"캔", 38},
	// }
	// sort.Sort(students)
	// fmt.Println(students)
}