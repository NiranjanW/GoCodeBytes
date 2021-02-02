package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	i, j := 10, 2701
	p := &i
	fmt.Println(&i, &j, *p)

	//*int in from of type then whole thing is a type , *p a variable a operaor the value dereference operator
	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p / 37
	fmt.Println(j)

	//go routine independant flow/work
	//gets its own stack

	a := 4
	//value semantics
	squareVal(a)

	//Modify the pointer semantics - give up immutability for efficency
	squareAdd(&a)
	fmt.Println(a)
	fmt.Println(initPerson())
	fmt.Printf("initPerson-->%p\n", initPerson())

}

func initPerson() *person {
	//p is copied to heap , Heap needs garage collection stck doesnt
	p := person{name: "Niran", age: 50}
	fmt.Printf("initPerson-->%p\n", &p)
	return &p
}

func squareVal(v int) {
	v *= v
	fmt.Println(&v, v)
}

func squareAdd(p *int) {
	*p *= *p
	fmt.Println(p, *p)

}
