package main

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"os"
	"time"
)

func updateTime(clock *widget.Label) {

	//上海时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	v, err := time.ParseInLocation(time.RFC3339, os.Args[1], loc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	t := gett(v)
	timed := string(fmt.Sprint(t.d))
	timeh := string(fmt.Sprint(t.h))
	timem := string(fmt.Sprint(t.m))
	times := string(fmt.Sprint(t.s))
	//timems := string(fmt.Sprint(t.ms))
	var timeall string = fmt.Sprint("距离", os.Args[2], "还有", timed, "天", timeh, "小时", timem, "分钟", times, "秒") //, timems, "毫秒")
	clock.SetText(timeall)
}

// 定义类型
type countdown struct {
	t int
	d int
	h int
	m int
	s int
	//ms int
}

func inputTime() time.Time {
	var endTime string
	var v time.Time
	fmt.Scanf("%d", &endTime)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	v, err := time.ParseInLocation(time.RFC3339, endTime, loc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return v

}

// 计算时间
func gett(t time.Time) countdown {

	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)
	//毫秒可能~~会炸~~或者**不准**
	//milliseconds := int(1000 / seconds)

	return countdown{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
		//ms: milliseconds,
	}
}

func main() {

	fmt.Println(inputTime())

}
