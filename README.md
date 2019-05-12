# X-Clock

A tiny executable binary used to display time information in system-tray.

## Build

- `go get github.com/prakashpandey/x-clock`
- `cd $GOPATH/src/github.com/prakashpandey/x-clock`
- `make build`
- `make run`

## Install

### Install for the current user

This will install binary to your `$GOPATH/bin`

- `make install`
- Now you can start the application by typing `x-clock` in your command line.

### Install application system wide

This will install application in `/usr/bin` and also make desktop entry.

- `chmod +x ./bin/install.sh`
- `./bin/install.sh`

## Screenshot

![X-Clock](assets/clock-screenshot.png)

## Based on

[getlantern/systray](https://github.com/getlantern/systray)

## License

Entire application is released with [MIT License](LICENSE) other than the third party icon image `assets/clock.png` which is
taken from [icons8.com](https://icons8.com/icon/63250/watch). If you want to use the same clock icon `assets/clock.png` for your use case, please check its license by visiting [icons8.com](https://icons8.com/icon/63250/watch)