package main

import (
	"fmt"
	"packbet/cards"
	"packbet/horse"
)

func main() {
	fmt.Println("Horse movements: ", horse.Movements(&horse.Field{"a", 1}))
	fmt.Println("======================")
	desk := cards.FillDesk(2)
	fmt.Println(desk)
	cards.Sorting(&desk)
	fmt.Println(desk)
}
