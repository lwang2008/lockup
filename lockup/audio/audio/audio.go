// Copyright 2021, Lucas Wang
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package audio

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	decrypt "github.com/lwang2008/lockup/lib/decrypt"
	"github.com/lwang2008/lockup/lockup/audio/audioControls"
)

func Display(_key, _imageFile, _audioFile string, _errMsg *canvas.Text, _window fyne.Window) {
	mainFile, err := decrypt.Decrypt(_key, _audioFile)
	if err != nil {
		_errMsg.Show()
		_errMsg.Text = err.Error()
	}
	setContent(mainFile, "", _window)
}

func setContent(_audioFile, _imageFile string, _window fyne.Window) {

	mediaControls := audioControls.GetAudioControls(_audioFile)
	_window.SetContent(mediaControls)
}
