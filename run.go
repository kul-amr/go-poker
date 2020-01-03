package main

import (
	"go-poker/poker"
	"log"
	"os"
)

func main() {
	//Collect arguments - hands to compare
	args := os.Args[1:]

	if len(args) == 2 {
		winHand1, windHand2 := poker.HandCompare(args[0], args[1])

		if winHand1 && !windHand2 {
			log.Println("First hand wins!")
		} else if windHand2 && !winHand1 {
			log.Println("Second hand wins!")
		} else if winHand1 && windHand2 {
			log.Println("It's a tie!")
		} else {
			log.Fatalln("Invalid hand/s - check the arguments passed")
		}
	} else {
		log.Fatalln("Invalid Arguments")
	}

}
