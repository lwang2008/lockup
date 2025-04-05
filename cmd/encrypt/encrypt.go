// Copyright 2021, Lucas Wang
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"image/color"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/lwang2008/lockup/lib/encrypt"
)

func main() {

	mainApp := app.New()
	mainWindow := mainApp.NewWindow("Encrypt file")
	mainWindow.Resize(fyne.Size{Width: 320, Height: 180})

	pname := widget.NewEntry()
	pname.SetPlaceHolder("Enter the name of your file")

	ename := widget.NewEntry()
	ename.SetPlaceHolder("Enter the name of your target file")

	key := widget.NewEntry()
	key.SetPlaceHolder("Type in the encryption key")

	messageCanvas := canvas.NewText("", theme.ForegroundColor())
	messageCanvas.TextStyle.Bold = true
	messageCanvas.Hide()

	submitButton := widget.NewButton("Submit", func() {

		ciphertext, err := encrypt.Encrypt(key.Text, pname.Text)
		if err != nil {
			messageCanvas.Show()
			messageCanvas.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
			messageCanvas.Text = "Error: " + err.Error()
		}

		err = os.WriteFile(ename.Text, ciphertext, 0777)
		if err != nil {
			messageCanvas.Show()
			messageCanvas.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255}
			messageCanvas.Text = "Error: " + err.Error()
		}

		messageCanvas.Show()
		messageCanvas.Text = "Success! You may now quit the app"
	})

	content := container.NewVBox(pname, ename, key, submitButton, messageCanvas)

	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}
