package Flash

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"io"

	//Potion "github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket"
	//Cyberball "github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket/Ballpocket"
	TM29 "github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket/TMpocket"

)

func Flash() {
	for {
		TM29.Cls()
		fmt.Println("Choose One")
		fmt.Println("1. Eat Me")
		fmt.Println("2. Drink Me")
		fmt.Println("3. Wake Up")
		res, err := TM29.StringPrompt("One or Two, What shall I do?")
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		switch res {
		case "1":
			response, err := TM29.FetchFromServer("/eatme")
			if err != nil {
				fmt.Println("Error fetching data:", err)
				continue
			}
			fmt.Println(response)
			err = TM29.ListAndSelectFiles("./EatMe")
			if err != nil {
				fmt.Println("Error selecting file:", err)
			}
		case "2":
			response, err := TM29.FetchFromServer("./drinkme")
			if err != nil {
				fmt.Println("Error fetching data:", err)
				continue
			}
			fmt.Println(response)
			err = TM29.ListAndSelectFiles("./DrinkMe")
			if err != nil {
				fmt.Println("Error selecting file:", err)
			}
		case "3":
			for {
				TM29.Cls()
				fmt.Println("Are You Sure?")
				fmt.Println("1. No or 2. Yes")
				resCheck, err := TM29.StringPrompt("Choose Wisely")
				if err != nil {
					fmt.Println("Error reading input:", err)
					continue
				}
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

func fetchFromServer(s string) (string, error) {
	resp, err := http.Get(s)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	body, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}

	return body, nil
}


func listAndSelectFiles(s string) (string, error) {
	files, err := os.ReadDir(s)
	if err != nil {
		return "", err
	}

	for i, file := range files {
		fmt.Printf("%d. %s\n", i+1, file.Name())
	}

	selection, err := TM29.stringPrompt("Select a file by number:")
	if err != nil {
		return "", err
	}
	index, err := strconv.Atoi(selection)
	if err != nil || index < 1 || index > len(files) {
		return "", fmt.Errorf("Invalid selection")
	}

	return files[index-1].Name(), nil
}
