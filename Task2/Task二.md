# Task二

#### 1.

```go
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

```



#### 2.

map:内存用 make 方法来分配。当 map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1。

slice:扩容是创建一个新的更大容量的数组，将原数组元素复制进来。

#### 3.

```go
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

```



#### 4.

若所有元素都可比较则可以直接进行比较struct1==struct2,struct1.a==struct2.a

反之不可以

#### 5.

```Go
  func main() {
      x := "hello!"
      for i := 0; i < len(x); i++ {
          x := x[i]
          if x != '!' {
              x := x + 'A' - 'a'
              fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
         }
     }
  }
```

第四行和第六行，x在for语句和if语句中重新声明了，覆盖了外层变量x