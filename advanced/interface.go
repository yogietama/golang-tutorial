// goplay.space/#C78gzH8CjD
package main

import "fmt"

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func main() {
	var r geometry
	r = rect{width: 5, height: 3}

	fmt.Println(r.area())
	fmt.Println(r.perim())
}
