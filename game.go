package main

import "fmt"

func main() {
	var choice int
	fmt.Println("Welcome to GoAdventure!")
	fmt.Println("You find yourself surrounded by craggy cliffs and dense fog. It's hard")
	fmt.Println("to tell apart the various paths that lie ahead... What will you do?")
	fmt.Println("1. Take the left path")
	fmt.Println("2. Take the right path")

	fmt.Scanf("%d", &choice)
	if (choice == 1) {
		leftPath()
	} else if (choice == 2) {
		rightPath()
	} else {
		fmt.Println("You hesitate, unsure of which path to take. The fog thickens") 
		fmt.Println("and you feel a chill run down your spine. Suddenly, you get")
		fmt.Println("blown off the mountain by a gust of wind. GAME OVER.")
	}
}

func leftPath() {
}

func rightPath() {
}