package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"

    HM5 "github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket/HMpocket"
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("HM5 Client")

    label := widget.NewLabel("Choose One")

    eatMeButton := widget.NewButton("Eat Me", func() {
        response, err := HM5.FetchFromServer("/eatme")
        if err != nil {
            label.SetText("Error: " + err.Error())
            return
        }
        label.SetText(response)
        // Additional logic for "Eat Me" can be added here
    })

    drinkMeButton := widget.NewButton("Drink Me", func() {
        response, err := HM5.FetchFromServer("/drinkme")
        if err != nil {
            label.SetText("Error: " + err.Error())
            return
        }
        label.SetText(response)
        // Additional logic for "Drink Me" can be added here
    })

    wakeUpButton := widget.NewButton("Wake Up", func() {
        // For "Wake Up", you might want to simply close the app or perform other actions
        myApp.Quit()
    })

    content := container.NewVBox(
        label,
        eatMeButton,
        drinkMeButton,
        wakeUpButton,
    )

    myWindow.SetContent(content)
    myWindow.ShowAndRun()
}
