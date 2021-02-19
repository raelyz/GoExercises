package main

func main() {
	customer1 := customer{
		firstName: "Micheal",
		lastName:  "Jordan",
		userName:  "MJ2020",
		passWord:  "1234567",
		email:     "MJ2020@gmail.com",
		phone:     1234567,
		address:   "18227 Capstan Greens Road Cornelius, NC 28031",
	}
	customer1.printAll()
	customer1.userCredentials()
	customer1.userAddress()
}
