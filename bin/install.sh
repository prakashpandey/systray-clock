app_name="x-clock"
app_desktop_file="x-clock.desktop"
printf "installing application: %s...\n" $app_name
sudo $(which go) build -o /usr/bin/$app_name
echo "making desktop entry..."
cp ./bin/$app_desktop_file ~/Desktop/$app_desktop_file
printf "\nIcon=$HOME/.x-clock/clock-desktop.png\n" >> ~/Desktop/$app_desktop_file
echo "installed successfully"