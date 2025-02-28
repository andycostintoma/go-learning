package main

import "fmt"

type Foo struct {
	x int
}

func main() {
	outer()
}

func outer() {
	f := &Foo{10} // f is a pointer to a Foo struct where x = 10
	inner1(f)
	fmt.Println(f.x) // prints 20
	inner2(f)
	fmt.Println(f.x) // still prints 20
	var g *Foo
	inner2(g)
	fmt.Println(g == nil) // prints true
}

func inner1(f *Foo) {
	f.x = 20 // Modifies the object that f points to
}

func inner2(f *Foo) {
	f = &Foo{30} // Reassigns f to a new Foo instance, but this does not affect the original f in outer
}
