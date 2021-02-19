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
	book1 := book{
		title: "Candle in the tomb",
		price: 20,
	}
	book2 := book{
		title: "Barney and Friends",
		price: 10,
	}
	merch1 := computerAccessories{
		title: "Razer BT earpiece",
		price: 159,
	}
	merch2 := computerAccessories{
		title: "Razer keyboard",
		price: 110,
	}
	merch3 := computerAccessories{
		title: "Logitech Mouse",
		price: 80,
	}
	slice := []merchandise{game1, game2, game3, book1, book2, merch1, merch2, merch3}

	list(slice)

}
