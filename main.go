package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Kiosk Clock")

	clock := canvas.NewText("", nil)
	clock.Alignment = fyne.TextAlignCenter
	clock.TextSize = 200

	updateClock := func() {
		clock.Text = time.Now().Format("15:04:05")
		clock.Refresh()
	}
	updateClock()

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			fyne.Do(updateClock)
		}
	}()

	w.SetContent(container.NewCenter(clock))
	w.SetFullScreen(true)
	w.ShowAndRun()
}
