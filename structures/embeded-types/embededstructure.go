package main
import "fmt"

type Person struct {
	name string
}

type Android struct {
	kind string
	Person
}

func (p Person) talk() {
	fmt.Println("My name is: ", p.name)
}

func main() {
	p := Person{"Susan"}
	p.talk()

	a := Android{"A11", p}
	a.talk()
}
