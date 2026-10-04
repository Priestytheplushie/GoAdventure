package main

import "fmt"

func main() {
	fmt.Println("GoAdventure")
	fmt.Println("1. Play")
	fmt.Println("2. Exit")

	var choice int
	fmt.Scanf("%d\n", &choice)

	if choice == 1 {
		start()
	} else if choice == 2 {
		fmt.Println("Exiting GoAdventure...")
	} else {
		fmt.Println("Invalid choice. Exiting GoAdventure...")
	}
}

func start() {
	var choice int
	fmt.Println("Welcome to GoAdventure!")
	fmt.Println("You find yourself surrounded by craggy cliffs and dense fog. It's hard")
	fmt.Println("to tell apart the various paths that lie ahead... What will you do?")
	fmt.Println("1. Take the left path")
	fmt.Println("2. Take the right path")

	fmt.Scanf("%d\n", &choice)
	if choice == 1 {
		leftPath()
	} else if choice == 2 {
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
	fmt.Scanf("%d\n", &choice)

	if choice == 1 {
		fmt.Println("You step into the dark cave, and stumble around. You hear a growl in the distance")
		fmt.Println("and start to get spooked. You realize that this may not have been the best idea, and")
		fmt.Println("you try to turn back, but it's too late, a bear leaps from behind a rock and eats your")
		fmt.Println("face off. GAME OVER.")
	} else if choice == 2 {
		fmt.Println("You continue on the path, but the fog grows thicker and you lose your way. As you")
		fmt.Println("wander aimlessly, you stumble and fall off the cliff. GAME OVER.")
	} else if choice == 3 {
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
	fmt.Scanf("%d\n", &choice)

	if choice == 1 {
		fmt.Println("You step into the warm cabin, the cozy vibe makes you feel at home. Alas, you")
		fmt.Println("are safe... YOU WIN!")
	} else {
		fmt.Println("You decide not to enter the cabin and continue on your way. As you walk away")
		fmt.Println("you start to feel a chill run down your spine, you hear distant noises from")
		fmt.Println("the fog, and you realize that you are not alone. Suddenly, a pack of wolves")
		fmt.Println("emerges from the fog and attacks you. You must decide quickly, do you:")
		fmt.Println("1. Run")
		fmt.Println("2. Fight")
		fmt.Scanf("%d\n", &choice)

		if choice == 1 {
			fmt.Println("You try to flee, but trip over a rock and fall. The wolves circle you, howling")
			fmt.Println("through the night, and you realize this is the end. You were never heard from again. GAME OVER.")
		} else if choice == 2 {
			fmt.Println("You stand your ground and fight the wolves. They look infuriated, but you think")
			fmt.Println("you can take them. What will you do?")
			fmt.Println("1. Use your fists")
			fmt.Println("2. Use a stick")
			fmt.Scanf("%d\n", &choice)

			if choice == 1 {
				fmt.Println("You throw a punch at the wolf, but the wolf just bites your hand off. You scream")
				fmt.Println("in agony, and the wolves circle in. You were never seen again... GAME OVER.")
			} else if choice == 2 {
				fmt.Println("You grab a nearby stick and the wolves chase it. You use the distraction to escape")
				fmt.Println("and find your way back to safety.")
				fmt.Println("As you continue on your journey, you start to get really cold. You feel a wave of")
				fmt.Println("regret for not entering the cabin, but you keep moving forward.")
				aftermath()
			}
		}
	}
}

func aftermath() {
	var choice int
	fmt.Println("As you wander, you start to hallucinate things, such as wild beasts and other animals")
	fmt.Println("that you swear are in front of you, but they are not. You let out a desperate cry")
	fmt.Println("and it attracts some distant creatures. What will you do?")
	fmt.Println("1. Run")
	fmt.Println("2. Hide")
	fmt.Println("3. Approach them")
	fmt.Scanf("%d\n", &choice)

	if choice == 1 {
		fmt.Println("You run away like a coward. The distant voice calls out to you, but you keep running.")
		fmt.Println("As you flee, you spot a cliff and it seems like the end of the line. You find nothing")
		fmt.Println("left to do, so you wait... and wait... and wait... but nobody comes. You were never seen again... GAME OVER.")
	} else if choice == 2 {
		fmt.Println("You duck behind a nearby rock and wait for the creature to pass, but it never does.")
		fmt.Println("It seems like it's waiting for you to make the first move.")
	} else if choice == 3 {
		fmt.Println("You approach the creatures, and they seem to be friendly. They offer you some food and water,")
		fmt.Println("and you accept. You feel a sense of relief and fulfillment. YOU WIN!")
	}
}