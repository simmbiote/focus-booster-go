package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Pomodoro Timer")

	label := widget.NewLabel("Ready")
	startButton := widget.NewButton("Start 25 minute focus", func() {
		go func() {
			totalSeconds := 25 * 60
			for i := totalSeconds; i >= 0; i-- {
				minutes := i / 60
				seconds := i % 60
				label.SetText(fmt.Sprintf("%02d:%02d", minutes, seconds))
				time.Sleep(1 * time.Second)
			}
			label.SetText("Time's up!")
		}()
	})

	myWindow.SetContent(container.NewVBox(
		label,
		startButton,
	))

	myWindow.ShowAndRun()
}
