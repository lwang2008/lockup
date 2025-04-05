// Copyright 2021, Lucas Wang
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package music

import (
	"fmt"
	"image"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	decrypt "github.com/lwang2008/lockup/lib/decrypt"
	"github.com/lwang2008/lockup/lockup/audio/audioControls"
)

func display(_key, _imageFile, _audioFile string, _errMsg *canvas.Text, _window fyne.Window) {

	audioFile, err := decrypt.Decrypt(_key, _audioFile)
	if err != nil {
		_errMsg.Show()
		_errMsg.Text = err.Error()
	}

	musicImage, err := decrypt.Decrypt(_key, _imageFile)
	if err == nil {
		setContent(audioFile, musicImage, _window)
	} else {
		_errMsg.Show()
		_errMsg.Text = err.Error()
		fmt.Print(err)

	}
}

func setContent(_audioFile, _imageFile string, _window fyne.Window) {

	reader := strings.NewReader(_imageFile)

	finalImg, _, sErr := image.Decode(reader)
	if sErr != nil {
		panic(sErr)
	}

	image := canvas.NewImageFromImage(finalImg)
	image.FillMode = canvas.ImageFillOriginal

	mediaControls := audioControls.GetAudioControls(_audioFile)
	finalContainer := container.NewVBox(image, mediaControls)

	_window.SetContent(finalContainer)
}
