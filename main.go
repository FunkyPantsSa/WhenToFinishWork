package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	//从现在到18:30的倒计时

	//endTime := time.Now().Add(18*time.Hour + 30*time.Minute)
	endTime := time.Now().Add(8 * time.Hour)
	formatted := endTime.Format("Time: 03:04:05")
	clock.SetText(formatted)

}

func main() {
	a := app.New()
	w := a.NewWindow("Clock")

	clock := widget.NewLabel("")
	updateTime(clock)

	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	w.ShowAndRun()
}
