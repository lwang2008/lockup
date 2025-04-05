package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	app := app.New()
	mainWin := app.NewWindow("View your NFT")

	topMessage := widget.NewLabel("Enter the link of the Lockup file")
	link := widget.NewEntry()
	link.SetPlaceHolder("Link")

	submit := widget.NewButton("Submit", func() {

		out, _ := os.Create("lockup.zip")

		resp, err := http.Get(link.Text)
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			panic(err)
		}

		exec.Command("unzip", "-j", "lockup.zip").Run()
		out.Close()
		resp.Body.Close()
		mainWin.Close()
		exec.Command("./lockup").Run()

	})

	container := container.NewVBox(topMessage, link, submit)
	mainWin.SetContent(container)
	mainWin.ShowAndRun()
}
