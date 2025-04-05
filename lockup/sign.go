package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/lwang2008/lockup/lib/ethRPC"
	"github.com/lwang2008/lockup/lockup/params"
)

func execute(_hash []byte, _hashData, _owner, _key, _imageFile, _audioFile string, _title *widget.Label, _errMsg *canvas.Text, _lockupWin fyne.Window) {
	signRequest := widget.NewEntry()
	signRequest.SetPlaceHolder(_hashData)
	signRequest.SetText("Sign this data " + _hashData)

	sigInput := widget.NewEntry()
	sigInput.SetPlaceHolder("Type Signature Here: ")

	submit := widget.NewButton("Submit", func() {
		if ethRPC.VerifySignature(string(sigInput.Text), _hash, _owner) {

			params.Display(_key, _imageFile, _audioFile, _errMsg, _lockupWin)

		} else {

			_errMsg.Show()
			_errMsg.Text = "ERROR: Malformed Signature"
		}
	})

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	port := os.Args[1]
	fmt.Println("Spawning server on http://localhost:" + port)

	websiteLink, err := url.Parse("http://localhost:" + port)
	if err != nil {
		panic(err)
	}

	signingLink := widget.NewHyperlink("link", websiteLink)

	http.HandleFunc("/getSignatureData", func(rw http.ResponseWriter, r *http.Request) {

		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			rw.WriteHeader(200)
			rw.Write([]byte("{\"value\":\"" + _hashData + "\"}"))
		} else {
			rw.WriteHeader(405)
		}
	})

	http.HandleFunc("/submitSignature", func(rw http.ResponseWriter, r *http.Request) {

		response, _ := io.ReadAll(r.Body)
		if ethRPC.VerifySignature(string(response), _hash, _owner) {

			rw.WriteHeader(200)
			params.Display(_key, _imageFile, _audioFile, _errMsg, _lockupWin)

		} else {
			rw.WriteHeader(400)
			_errMsg.Show()
			_errMsg.Text = "ERROR: Malformed Signature"
		}
	})

	go http.ListenAndServe(":"+fmt.Sprint(port), nil)

	topContent := container.NewHBox(_title, signingLink)
	content := container.NewVBox(signRequest, topContent, sigInput, submit, _errMsg)
	_lockupWin.SetContent(content)
}
