package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	//从现在到18:30的倒计时
	now := time.Now()
	//target := time.Now().Add(18*time.Hour + 30*time.Minute)
	//endTime := time.Now().Add(8 * time.Hour)
	target := time.Date(now.Year(), now.Month(), now.Day(), 18, 00, 0, 0, now.Location())

	// 如果当前时间晚于目标时间，则目标时间延后一天
	if now.After(target) {
		target = target.AddDate(0, 0, 1)
	}

	// 计算倒计时时间
	var duration = target.Sub(now).String()
	//formatted := strconv.FormatInt(int64(duration), 10)

	clock.SetText(duration)
	//fmt.Println(duration)

}

func main() {
	a := app.New()
	w := a.NewWindow("Clock")

	content_time := widget.NewLabel("")
	updateTime(content_time)

	w.SetContent(content_time)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(content_time)

		}
	}()
	//content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

	/*	content := widget.NewButton("setting", func() {
		log.Println("tapped")
	})*/

	//content := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
	//	log.Println("tapped home")
	//})
	//content2 := container.New(layout.NewHBoxLayout(), content_time, content, layout.NewSpacer())
	w.SetContent(content_time)
	w.ShowAndRun()

}
