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
	var choice int
	fmt.Println("You take the right path, and follow the winding road, the fog begins to clear")
	fmt.Println("and you find a lone cabin with smoke rising from the chimney. You knock on the door")
	fmt.Println("but no one answers. Will you enter?")
	fmt.Println("1. Yes")
	fmt.Println("2. No")
	fmt.Scanf("%d", &choice)
	if (choice == 1) {
		fmt.Println("You step into the warm cabin, the cozy vibed makes you feel at home. Alas, you")
		fmt.Println("are safe... YOU WIN!")
	} else {
		fmt.Println("You decide not to enter the cabin and continue on your way. As you walk away")
		fmt.Println("you start to feel a chill run down your spine, you hear distant noises from")
		fmt.Println("the fog, and you realize that you are not alone. Suddenly, a pack of wolves")
		fmt.Println("emerges from the fog and attacks you. You must decide quickly, do you:")
		fmt.Println("1. Run")
		fmt.Println("2. Fight")
		fmt.Scanf("%d", &choice)
		if (choice == 1) {
			fmt.Println("You try to flee, but trip over a rock and fall, the wolves circle you, howling")
			fmt.Println("through the nightand you realize that this is the end. You we're never heard from again. GAME OVER.")
		}
		else if (choice == 2) {
			fmt.Println("You stand your ground, and fight the wolves, they look infuriated, but you think")
			fmt.Println("you can take them. What will you do?")
			fmt.Println("1. Use your fists")
			fmt.Println("2. Use a stick")
			fmt.Scanf("%d", &choice)
			if (choice == 1) {
				fmt.Println("You throw a punch at the wolf, but the wolf just bites your hand off. You scream")
				fmt.Println("in agony, and the wolves circle in. You were never seen again... GAME OVER.")
			}
			else if (choice == 2) {
				fmt.Println("You toss a nearby stick and the wolves dig it! They chase it around and you take")
				fmt.Println("the oppertunity to run away. You escape the wolves and find your way back to safety.")
				fmt.Println("As you continune on your journey, you start to get really cold. You feel a wave of")
				fmt.Println("regret for not entering the cabin, but you keep moving forward.")
				aftermath()
			}
		}
	}
}

func aftermath() {
	var choice int
	fmt.Println("As you wonder, you start to hallutionate things, such as wild beasts, and other animals who you")
	fmt.Println("swear are in fromnt of you, but theyre not. You let out a desperate cry, and it attracts some")
	fmt.Println("distant creatures. What will you do?")
	fmt.Println("1. Run")
	fmt.Println("2. Hide")
	fmt.Println("3. Approach them")
	fmt.Scanf("%d", &choice)
	if (choice == 1) {
		fmt.Println("You run away like a coward, the distant voice calls out to you, but you keep on running, what")
		fmt.Println("a loser. As you flee you spot a cliff, and it seems like the end of the line. As you look aroundc")
		fmt.Println("you find nothing left to do. So you just wait... and wait... and wait... but nobody came, and nothing")
		fmt.Println("happened. You were never seen again... GAME OVER.")
}
