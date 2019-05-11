package main

import (
	"fmt"
	"io/ioutil"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(run, exit)
}

func run() {
	setIcon()
	systray.SetTitle("x-clock")
	systray.SetTooltip("Clock application")
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
