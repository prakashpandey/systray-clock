package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/getlantern/systray"
)

const (
	appName  = "x-clock"
	tzIndia  = "Asia/Kolkata"
	tzAusSyd = "Australia/Sydney"
)

func main() {
	systray.Run(run, exit)
}

func run() {
	setIcon()
	go func() {
		for {
			title := fmt.Sprintf("IN-%s | SYD-%s", getTime(tzIndia), getTime(tzAusSyd))
			systray.SetTitle(title)
			time.Sleep(1 * time.Second)
		}
	}()
	systray.SetTooltip(appName)
}

func getTime(timezone string) string {
	loc, _ := time.LoadLocation(timezone)
	h, m, s := time.Now().In(loc).Clock()
	return fmt.Sprintf("%s:%s:%s", appendZeroIfSingleDigitInteger(h), appendZeroIfSingleDigitInteger(m), appendZeroIfSingleDigitInteger(s))
}

func appendZeroIfSingleDigitInteger(i int) string {
	if 0 <= i && i < 10 {
		return fmt.Sprintf("%s%d", "0", i)
	}
	return fmt.Sprintf("%d", i)
}

func setIcon() {
	file := "assets/clock.ico"
	icon, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("error reading app icon file: %s : %s", file, err)
	}
	systray.SetIcon(icon)
}

func exit() {

}
