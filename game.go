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
	var choice int 
	fmt.Println("You take the left path and push through the fog, after a few long")
	fmt.Println("steps you discover a hidden cave. Shivering from the cold, you have")
	fmt.Println("a choice to make:")
	fmt.Println("1. Enter the cave")
	fmt.Println("2. Continue on the path")
	fmt.Println("3. Turn back and take the right path")
	fmt.Scanf("%d", &choice)
	if (choice == 1) {
		fmt.Println("You step into the dark cave, and stumble around. You hear a growl in the distance")
		Fmt.Println("and start to get spooked. You realize that this may not have been the best idea, and")
		fmt.Println("you try to turn back, but it's too late, a bear leaps from behind a rock and eats your")
		fmt.Println("face off. GAME OVER.")
	} else if (choice == 2) {
		fmt.Println("You continue on the path, but the fog grows thicker and you lose your way. As you")
		fmt.Println("wander aimlessly, you stumble and fall off the cliff. GAME OVER.")
	} else if (choice == 3) {
		fmt.Println("You turn back and take the right path.")
		rightPath()
	} else {
		fmt.Println("You hesitate, unsure of which path to take. The fog thickens") 
		fmt.Println("and you feel a chill run down your spine. Suddenly, you get")
		fmt.Println("blown off the mountain by a gust of wind. GAME OVER.")
	}
}

func rightPath() {
	fmt.Println("You take the right path, and follow the winding road, the fog begins to clear")
	fmt.Println("and you find a lone cabin with smoke rising from the chimney. You knock on the door")
	fmt.Println("but no one answers. Will you enter?")
}