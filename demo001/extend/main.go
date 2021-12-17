package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name  string
	Age   int
	Child *Person
}

func print(p *Person, pname string) {
	if p == nil {
		return
	}
	if p.Child != nil {
		fmt.Printf("%s name %s age %d c.n %s c.a %d \n", pname, p.Name, p.Age, p.Child.Name, p.Child.Age)
	} else {
		fmt.Printf("%s name %s age %d \n", pname, p.Name, p.Age)
	}
}

func main() {
	p1 := &Person{Name: "p1", Age: 10}
	print(p1, "p1")

	p2 := p1
	print(p2, "p2")

	p2.Child = &Person{Name: "c1", Age: 1}
	fmt.Printf("type of %v %v \n", reflect.TypeOf(p1.Child), reflect.TypeOf(*p1.Child))
	print(p2, "\np2")
	print(p1, "p1")

	p2 = p2.Child
	print(p2, "\np2")
	print(p1, "p1")

	print(p1.Child, "\np1.c")
	print(p2.Child, "p2.c")

	p2.Child = &Person{Name: "c11", Age: 11}
	print(p2, "\np2")
	print(p1, "p1")

	print(p1.Child, "\np1.c")
	print(p1.Child.Child, "\np1.c.c")
	print(p2.Child, "p2.c")
}
