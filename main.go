package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"time"

	"github.com/getlantern/systray"
)

const (
	appName        = "x-clock"
	tzIndia        = "Asia/Kolkata"
	tzAusSyd       = "Australia/Sydney"
	appHomeDirName = ".x-clock"
	iconFileName   = "clock.png"
	remoteIconURL  = "https://raw.githubusercontent.com/prakashpandey/x-clock/master/assets/clock.png"
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

func setIcon() error {
	appHome, err := getAppHome()
	if err != nil {
		return err
	}
	filePath := fmt.Sprintf("%s/%s", appHome, iconFileName)
	if !fileExist(filePath) {
		if _, err := downloadFile(remoteIconURL, appHome, iconFileName); err != nil {
			return err
		}
	}
	icon, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("error reading app icon file: %s : %s", filePath, err)
	}
	systray.SetIcon(icon)
	return nil
}

func fileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
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

func downloadFile(url, targetDirectory, targetFileName string) (*os.File, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	} else if !statusOk(resp.StatusCode) {
		return nil, fmt.Errorf("failed to download file from url: %s, status code: %d", url, resp.StatusCode)
	}
	defer resp.Body.Close()
	// create dir
	if err := os.MkdirAll(targetDirectory, 0755); err != nil {
		return nil, err
	}
	targetPath := fmt.Sprintf("%s/%s", targetDirectory, targetFileName)
	file, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(file, resp.Body); err != nil {
		return nil, err
	}
	return file, nil
}

func statusOk(code int) bool {
	return 200 <= code && code < 400
}

func getAppHome() (string, error) {
	userHome, err := getUserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", userHome, appHomeDirName), nil
}

func getUserHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}
