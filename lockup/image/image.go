// Copyright 2021, Lucas Wang
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package image

import (
	"image"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	decrypt "github.com/lwang2008/lockup/lib/decrypt"
)

func Display(_key, _imageFile, _audioFile string, _errMsg *canvas.Text, _window fyne.Window) {
	mainFile, err := decrypt.Decrypt(_key, _imageFile)
	if err != nil {
		_errMsg.Show()
		_errMsg.Text = err.Error()
	}
	setContent("", mainFile, _window)
}
func setContent(_audioFile, _imageFile string, _window fyne.Window) {

	reader := strings.NewReader(_imageFile)

	finalImg, _, sErr := image.Decode(reader)
	if sErr != nil {
		panic(sErr)
	}

	image := canvas.NewImageFromImage(finalImg)

	image.FillMode = canvas.ImageFillOriginal

	_window.SetContent(image)
}
