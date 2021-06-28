package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	peterbilt := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}

	crown := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: true,
	}

	fmt.Println(peterbilt)
	fmt.Println(crown)

	fmt.Println(crown.color)
}
