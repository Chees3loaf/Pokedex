package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	HM5 "github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket/HMpocket"
)

// main initializes and runs the Fyne application.
func main() {
	// Create a new Fyne application.
	myApp := app.New()

	// Create a new window with the title "HM5 Client".
	myWindow := myApp.NewWindow("HM5 Client")

	// Create a label widget to display instructions or responses.
	label := widget.NewLabel("Choose One")

	// Create a button for the "Eat Me" option.
	eatMeButton := widget.NewButton("Eat Me", func() {
		// When the button is clicked, fetch data from the server.
		response, err := HM5.FetchFromServer("/eatme")
		if err != nil {
			// If there's an error, display it on the label.
			label.SetText("Error: " + err.Error())
			return
		}
		// Display the server's response on the label.
		label.SetText(response)
		// Additional logic for "Eat Me" can be added here.
	})

	// Create a button for the "Drink Me" option.
	drinkMeButton := widget.NewButton("Drink Me", func() {
		// Similar to "Eat Me", fetch data for "Drink Me".
		response, err := HM5.FetchFromServer("/drinkme")
		if err != nil {
			label.SetText("Error: " + err.Error())
			return
		}
		label.SetText(response)
		// Additional logic for "Drink Me" can be added here.
	})

	// Create a button for the "Wake Up" option.
	wakeUpButton := widget.NewButton("Wake Up", func() {
		// This button will close the application.
		myApp.Quit()
	})

	// Arrange the label and buttons vertically in the window.
	content := container.NewVBox(
		label,
		eatMeButton,
		drinkMeButton,
		wakeUpButton,
	)

	// Set the content of the window to the arranged widgets.
	myWindow.SetContent(content)

	// Start the application and display the window.
	myWindow.ShowAndRun()
}
