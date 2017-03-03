package main

import (
	"errors"
	"fmt"
)

 Persion struct{
	name stirng,
	pwd string,
	ape int
}

var a, b = 25, "abc"

func test(a int, b int) {
	if a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
	fmt.Println(a + b)
}

func strtest(s string) {
	c := []byte(s)
	c[0] = 'd'
	f := string(c)
	fmt.Println(f)
}

func errtest() {
	err := errors.New("有错误")
	if err != nil {
		fmt.Println(err)
	}
	err = errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//	errtest()
	//	test(12, 13)
	//	fmt.Println("test")
	//	fmt.Println(a, b)
	//	strtest(b)

	a := funtest(15, 17)
	fmt.Println(a)
	 for _, v := range map{
        fmt.Println("map's val:", v)
    }

}
