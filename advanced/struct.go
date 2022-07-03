package main

import "fmt"

type MyBox struct {
	Length float64
	Width  float64
	Height float64
	color  string
}

func (mb MyBox) Volume() float64 {
	return mb.Length * mb.Width * mb.Height
}

func (mb *MyBox) setColor(color string) {
	mb.color = color
}

// incorrect use of receiver
func (mb MyBox) SetLength(l float64) {
	mb.Length = l
}

func main() {
	myBox := MyBox{
		Length: 12.5,
		Width:  5.3,
		Height: 20,
	}

	fmt.Println(myBox)                 // {12.5 5.3 20 }
	fmt.Println("Width:", myBox.Width) // 12.5
	myBox.Width = 7.5
	fmt.Println("Width:", myBox.Width)

	fmt.Println(myBox.Volume()) // 12.5 * 7.5 * 20
	myBox.setColor("red")
	myBox.SetLength(100) //didnt work, length didnt change to 100
	fmt.Println(myBox)   // {12.5 7.5 20 red}

	// myBox with Pointer
	myBoxPtr := &MyBox{
		Length: 2.5,
		Width:  0.2,
		Height: 17,
	}

	fmt.Println("==== My Box Pointer ====")
	fmt.Println(myBoxPtr)
	myBoxPtr.setColor("blue")
	fmt.Println(myBoxPtr.color) // blue

	// does not work
	myBox.SetLength(989)
	fmt.Println(myBox.Length) // didnt chate to 989

	fmt.Println(myBoxPtr)

}
