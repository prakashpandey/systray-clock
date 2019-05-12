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
			title := fmt.Sprintf("IN %s | SYD %s", getTime(tzIndia), getTime(tzAusSyd))
			systray.SetTitle(title)
			time.Sleep(1 * time.Second)
		}
	}()
	systray.SetTooltip(appName)
}

func getTime(timezone string) string {
	loc, _ := time.LoadLocation(timezone)
	h, m, s := time.Now().In(loc).Clock()
	h, cy, _ := normalizeTo12Hour(h) // normalize 24 hr format to 12 hours
	return fmt.Sprintf("%s:%s:%s %s", appendZeroIfSingleDigitInteger(h), appendZeroIfSingleDigitInteger(m), appendZeroIfSingleDigitInteger(s), cy)
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

type cycle string

const (
	am cycle = "AM"
	pm cycle = "PM"
)

// normalizeTo12Hour converts 24 hours to 12 hours format
// with cycle details such as 'AM', 'PM'.
func normalizeTo12Hour(h int) (int, cycle, error) {
	if h < 0 || h >= 24 {
		return 0, am, fmt.Errorf("h: %d should be between 0-24 where 0 is included and 24 is excluded", h)
	}
	if h < 12 {
		if h == 0 {
			return 12, am, nil
		}
		return h, am, nil
	} else {
		if h == 12 {
			return h, pm, nil
		}
		return h - 12, pm, nil
	}
}

func exit() {
	fmt.Println("x-clock exiting ...")
}
