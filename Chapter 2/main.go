package main

import "fmt"

type Sayian struct {
	name  string
	power int
}

func main() {
	//BasicTypes()
	PointersSyian()
}

// Basic Type declarations
func BasicTypes() {

	goku := Sayian{
		name:  "Goku",
		power: 9000,
	}

	gokuTwo := Sayian{
		"Goku",
		9000,
	}

	var gokuThree = Sayian{name: "GokuThree", power: 9001}

	fmt.Println(goku)
	fmt.Println(gokuTwo)
	fmt.Print(gokuThree)

}

//Helper Function
func Super(s *Sayian) (*Sayian) {
	s.power += 10000
	return s
}

// Trying to understand how an address and value are used with pointers and objects

// & Address operator
// * Means Address to type of X
func PointersSyian() {

	goku := &Sayian{name: "Goku", power: 9000}
	goku = Super(goku)

	fmt.Println(goku)

}
