package main

import (
	"fmt"
	"os/exec"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	setup := app.New()
	mainWin := setup.NewWindow("Let's get started!")
	formatLabel := widget.NewLabel("Select the format of your work.")
	var fomatChoice string
	formatOptions := widget.NewRadioGroup([]string{"Audio", "Music", "Image"}, func(s string) { fomatChoice = s })
	formatBox := container.NewVBox(formatLabel, formatOptions)

	var viewChoice string
	viewLabel := widget.NewLabel("Select the signing method of the application.")
	viewOptions := widget.NewRadioGroup([]string{"Web", "Local"}, func(s string) { viewChoice = s })
	viewBox := container.NewVBox(viewLabel, viewOptions)

	submitButton := widget.NewButton("Submit", func() {
		out, err := exec.Command("mkdir", "src").Output()
		if err != nil {
			fmt.Println(out)
			panic(err)
		}

		out, err = exec.Command("cp", "params.go", "src").Output()
		if err != nil {
			fmt.Println(out)
			panic(err)
		}

		out, err = exec.Command("cp", "templates/base/lockup.go", "src").Output()
		if err != nil {
			fmt.Println(out)
			panic(err)
		}
		if fomatChoice == "Audio" || fomatChoice == "Music" {
			out, err := exec.Command("cp", "templates/audio/audioControls/audioControls.go", "src").Output()
			if err != nil {
				fmt.Println(out)
				panic(err)
			}

			if fomatChoice == "Audio" {
				out, err := exec.Command("cp", "templates/audio/audio/audio.go", "src").Output()
				if err != nil {
					fmt.Println(out)
					panic(err)
				}
			}

			if fomatChoice == "Music" {
				out, err := exec.Command("cp", "templates/audio/music/music.go", "src").Output()
				if err != nil {
					fmt.Println(out)
					panic(err)
				}
			}
		}

		if fomatChoice == "Image" {
			out, err := exec.Command("cp", "templates/image/image.go", "src").Output()
			if err != nil {
				fmt.Println(out)
				panic(err)
			}
		}

		if viewChoice == "Web" {
			out, err := exec.Command("cp", "-r", "templates/web/public", "src").Output()
			if err != nil {
				fmt.Println(out)
				panic(err)
			}
			out, err = exec.Command("cp", "templates/web/web.go", "src").Output()
			if err != nil {
				fmt.Println(out)
				panic(err)
			}
		}

		if viewChoice == "Local" {
			out, err := exec.Command("cp", "templates/local/local.go", "src").Output()
			if err != nil {
				fmt.Println(out)
				panic(err)
			}
		}

		out, err = exec.Command("go", "build", "-o", "build", "./src").Output()
		if err != nil {
			fmt.Println(out)
			panic(err)
		}
	})
	mainWin.SetContent(container.NewVBox(container.NewHBox(formatBox, viewBox), submitButton))
	mainWin.ShowAndRun()

}
