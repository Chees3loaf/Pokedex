package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	HM5 "path/to/your/HM5/package"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("HM5 Client")

	label := widget.NewLabel("Choose One")
	eatMeButton := widget.NewButton("Eat Me", func() {
		// Handle "Eat Me" logic
	})
	drinkMeButton := widget.NewButton("Drink Me", func() {
		// Handle "Drink Me" logic
	})
	wakeUpButton := widget.NewButton("Wake Up", func() {
		// Handle "Wake Up" logic
	})

	content := container.NewVBox(
		label,
		eatMeButton,
		drinkMeButton,
		wakeUpButton,
		// Add more widgets as needed
	)

	myWindow.SetContent(content)

	myWindow.ShowAndRun()

	eatMeButton = widget.NewButton("Eat Me", func() {
		response, err := HM5.FetchFromServer("/eatme")
		if err != nil {
			label.SetText("Error: " + err.Error())
			return
		}
		label.SetText(response)
		// Additional logic to handle file selection, etc.
	})
}
