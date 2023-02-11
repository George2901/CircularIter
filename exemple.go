package main

import (
	"main/circular_str"
)

func main() {
	circle := circular_str.NewCircle()
	for i := 0; i <= 20; i++ {
		circle.Append("pm")
	}
	for _, i := range circle.GerElements() {
		println(i)
	}

}
