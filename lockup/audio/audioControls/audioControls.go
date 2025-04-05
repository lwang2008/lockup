// Copyright 2021, Lucas Wang
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package audioControls

import (
	"io"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func GetAudioControls(_audioFile string) *fyne.Container {
	audioReader := io.NopCloser(strings.NewReader(_audioFile))

	streamerPre, format, err := mp3.Decode(audioReader)
	if err != nil {
		panic(err)
	}

	buffer := beep.NewBuffer(format)
	buffer.Append(streamerPre)
	streamerPre.Close()

	streamer := buffer.Streamer(0, buffer.Len())

	sr := format.SampleRate

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: true}

	speaker.Init(sr, sr.N(time.Second/10))
	speaker.Play(ctrl)

	prevButton := widget.NewButtonWithIcon("", theme.MediaFastRewindIcon(), func() {

		speaker.Lock()
		newPos := streamer.Position()
		newPos -= sr.N(time.Second)

		if newPos < 0 {
			newPos = 0
		}
		if newPos >= streamer.Len() {
			newPos = streamer.Len() - 5
		}

		if err := streamer.Seek(newPos); err != nil {
			panic(err)
		}

		speaker.Unlock()

	})

	playButton := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), func() {
		ctrl.Paused = !ctrl.Paused

	})

	nextButton := widget.NewButtonWithIcon("", theme.MediaFastForwardIcon(), func() {
		speaker.Lock()
		newPos := streamer.Position()
		newPos -= sr.N(time.Second)

		if newPos < 0 {
			newPos = 0
		}
		if newPos >= streamer.Len() {
			newPos = streamer.Len() + 5
		}

		if err := streamer.Seek(newPos); err != nil {
			panic(err)
		}

		speaker.Unlock()
	})

	mediaControls := container.NewHBox(layout.NewSpacer(), prevButton, playButton, nextButton, layout.NewSpacer())

	return mediaControls

}
