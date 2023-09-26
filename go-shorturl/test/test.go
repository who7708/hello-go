package main

import "fmt"

type Mutable struct {
	a int
	b int
}

func (m Mutable) StayTheSame() {
	m.a = 5
	m.b = 7
}

func (m *Mutable) Mutatable() {
	m.a = 5
	m.b = 7
}

func main() {
	m := &Mutable{a: 0, b: 0}
	fmt.Println(m)
	fmt.Println(&m)

	m.StayTheSame()
	fmt.Println(m)
	fmt.Println(&m)

	m.Mutatable()
	fmt.Println(m)
	fmt.Println(&m)
}
