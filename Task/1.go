package main

import "fmt"

func f() func() int {
	a := 0
	return func() int {
		a++
		return a
	}

}
func main() {
	b := f()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}
