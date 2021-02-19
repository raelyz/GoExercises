package main

func main() {
	game1 := game{
		title: "Minecraft",
		price: 5,
	}
	game2 := game{
		title: "World of Warcraft",
		price: 19,
	}
	game3 := game{
		title: "Elite Dangerous",
		price: 54,
	}
	slice := []game{game1, game2, game3}

	list(slice)

}
