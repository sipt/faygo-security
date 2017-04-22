package main

import "fmt"

type ITest interface {
	Eat()
}

type Test struct {
}

func (t *Test) Eat() {
	fmt.Println("eating...")
}

func main() {
	t := &Test{}
	var temp interface{}
	temp = t
	_, ok := temp.(ITest)
	fmt.Println(ok)
}
