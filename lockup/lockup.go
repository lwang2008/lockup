// Copyright 2021, Lucas Wang
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/lwang2008/lockup/lib/ethRPC"
	"github.com/lwang2008/lockup/lockup/params"
)

func main() {

	lockup := app.New()
	lockupWin := lockup.NewWindow(params.NftTitle)

	title := widget.NewLabel("Sign your data online using this")

	errMsg := canvas.NewText("", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	errMsg.TextStyle.Bold = true
	errMsg.Hide()

	hashData, hash, owner := ethRPC.VerifyTokenOwner(params.TokenId, params.ContractAddress, params.RpcURL)

	execute(hash, hashData, owner, params.Key, params.ImageFile, params.AudioFile, title, errMsg, lockupWin)

	lockupWin.ShowAndRun()

}
