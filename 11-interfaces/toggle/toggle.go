package main

import "fmt"

type Switch bool

func (s *Switch) toggle() {
	*s = !*s
	fmt.Println(*s)
}

type Toggleable interface {
	toggle()
}

func main() {
	s := Switch(true)

	var t Toggleable = &s
	t.toggle()
	t.toggle()
	t.toggle()
	t.toggle()
	t.toggle()
}
