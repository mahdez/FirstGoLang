package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
	for i := 1; i < 9; i++ {
		println(i)
	}
	println(printSchoolAddress())
	println(getSchoolAddress())
	TestArray()
}

func printSchoolAddress() string {

	return "กรุงเทพ"
}

func getSchoolAddress() (int, string) {
	code := 1993
	address := "กรุงเทพ"
	return code, address
}

func TestArray() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	a[1] = 50
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
}
