package main

import (
	"bufio"
	"net/http"
	"strconv"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/widget/form"
)

var (

	gameBoy 	*app.App
	redBlue		fyne.Window
	LeaderBoard *widget.Label
	urlEntry	*widget.Entry
	pokeBall	*widget.Button
)

func main() {
	gameBoy = app.New()
	redBlue = gameBoy.NewWindow("")

}