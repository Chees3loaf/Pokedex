package Flash

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	//Potion "github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket"
	//Cyberball "github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket/Ballpocket"
	TM29 "github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket/TMpocket"
)

// Flash is the main interactive loop of the application.
func Flash() {
	for {
		// Clear the screen before showing the menu.
		TM29.Cls()

		// Display the main menu options to the user.
		fmt.Println("Choose One")
		fmt.Println("1. Eat Me")
		fmt.Println("2. Drink Me")
		fmt.Println("3. Wake Up")

		// Prompt the user for a choice and handle any input errors.
		res, err := TM29.StringPrompt("One or Two, What shall I do?")
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue // Continue to the next iteration of the loop on error.
		}

		// Handle the user's choice using a switch statement.
		switch res {
		case "1":
			// Fetch data from the server for the "Eat Me" option.
			response, err := TM29.FetchFromServer("/eatme")
			if err != nil {
				fmt.Println("Error fetching data:", err)
				continue // Continue to the next iteration of the loop on error.
			}
			fmt.Println(response)

			// Allow the user to select a file after fetching data.
			selectedFile, err := TM29.ListAndSelectFiles("./EatMe")
			if err != nil {
				fmt.Println("Error selecting file:", err)
				continue // Continue to the next iteration of the loop on error.
			}
			fmt.Println("Selected file:", selectedFile)

		case "2":
			// Similar handling for the "Drink Me" option.
			response, err := TM29.FetchFromServer("./drinkme")
			if err != nil {
				fmt.Println("Error fetching data:", err)
				continue
			}
			fmt.Println(response)

			selectedFile, err := TM29.ListAndSelectFiles("./DrinkMe")
			if err != nil {
				fmt.Println("Error selecting file:", err)
				continue
			}
			fmt.Println("Selected file:", selectedFile)

		case "3":
			// A nested loop for the "Wake Up" option to confirm the user's intention to exit.
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
					break // Exit the nested loop, returning to the main menu.
				} else if resCheck == "2" {
					fmt.Println("Goodbye Friend!")
					os.Exit(0) // Exit the application.
				} else {
					fmt.Println("Please choose 1 or 2.")
					time.Sleep(2 * time.Second) // Pause before continuing the loop.
				}
			}

		default:
			// Handle any input that is not 1, 2, or 3.
			fmt.Println("1, 2 or 3, it's all the same to me...")
			time.Sleep(3 * time.Second) // Pause before continuing the loop.
		}
	}
}

// FetchFromServer makes an HTTP GET request to the specified URL and returns the response body.
// It returns the response as a string and any error encountered during the process.
func FetchFromServer(s string) (string, error) {
	// Perform an HTTP GET request to the URL specified by the string 's'.
	resp, err := http.Get(s)
	if err != nil {
		// If there's an error in making the request, return an empty string and the error.
		return "", err
	}
	// Ensure the response body is closed after the function exits.
	defer resp.Body.Close()

	// Create a new buffered reader to read the response body.
	reader := bufio.NewReader(resp.Body)
	// Read the response body up to the first newline character.
	// This assumes the response ends at the first newline.
	body, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		// If there's an error reading the response body (and it's not an EOF error),
		// return an empty string and the error.
		return "", err
	}

	// Return the response body as a string and nil for the error.
	return body, nil
}

// Lists files in a given directory and allows the user to select one.
// It returns the name of the selected file and any error encountered.
func listAndSelectFiles(s string) (string, error) {
	// Read the directory specified by the string 's'.
	files, err := os.ReadDir(s)
	if err != nil {
		// Return an empty string and the error if the directory cannot be read.
		return "", err
	}

	// Iterate over the files and print their names.
	for i, file := range files {
		fmt.Printf("%d. %s\n", i+1, file.Name())
	}

	// Prompt the user to select a file by number.
	selection, err := TM29.StringPrompt("Select a file by number:")
	if err != nil {
		// Return an empty string and the error if there's an issue with the input.
		return "", err
	}

	// Convert the user's input (a string) to an integer.
	index, err := strconv.Atoi(selection)
	if err != nil || index < 1 || index > len(files) {
		// Return an error if the input is not a valid number or is out of range.
		return "", fmt.Errorf("Invalid selection")
	}

	// Return the name of the selected file. The -1 adjusts the 1-based user input to the 0-based slice index.
	return files[index-1].Name(), nil
}
