package main

import "fmt"

func main() {
	a := "hello,世界"
	fmt.Println(len(a))
	b := 0
	for i, _ := range a {
		fmt.Println(i)
		b++
	}
	fmt.Println(b)

}
