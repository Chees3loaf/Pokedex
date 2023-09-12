package flash

import (

	"fmt"
	"os"
	"time"

	Potion "github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket"
	TM29 "github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket/TMPocket"
)


func flash () {

	for {
		cls()
		fmt.Println("Choose One")
		fmt.Println("1. Eat Me")
		fmt.Println("2. Drink Me")
		fmt.Println("3. Wake Up")
		res := stringPrompt("One or Two, What shall I do?")

		switch res {
		case "1":
			response := fetchFromServer("/eatme")
			fmt.Println(response)
			listAndSelectFiles("./EatMe")
		case "2":
			response := fetchFromServer("/drinkme")
			fmt.Println(response)
			listAndSelectFiles("./DrinkMe")
		case "3":
			for {
				cls()
				fmt.Println("Are You Sure?")
				fmt.Println("1. No or 2. Yes")
				resCheck := stringPrompt("Choose Wisely")
				if resCheck == "1" {
					break
				} else if resCheck == "2" {
					fmt.Println("Goodbye Friend!")
					os.Exit(0)
				} else {
					fmt.Println("Please choose 1 or 2.")
					time.Sleep(2 * time.Second)
				}
			}
		default:
			fmt.Println("1, 2 or 3, it's all the same to me...")
			time.Sleep(3 * time.Second)
		}
	}
}

func stringPrompt(s string) {
	panic("unimplemented")
}