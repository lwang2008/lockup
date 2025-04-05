package params

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	// Uncomment your preferred format if either audio or music and comment the image

	format "github.com/lwang2008/lockup/lockup/image" // IMAGE
	// format "github.com/lwang2008/lockup/lockup/templates/audio/audio" // AUDIO
	// format "github.com/lwang2008/lockup/lockup/templates/audio/music" // MUSIC
)

const (
	// START USER CHANGES

	// Token ID of the specific NFT respective to its contract address
	TokenId string = "1"

	// Smart contract address of the NFT
	ContractAddress string = "0x0000000000000000000000000000000000000000"

	// Key used to decrypt files
	Key string = "0000"

	// URL endpoint of the JSON RPC server
	RpcURL string = "http://exampleRPC.org/myAccessToken"

	// Title of the NFT
	NftTitle string = "Lockup"

	// Name of image encrypted file
	ImageFile string = "nft.bin"

	// Name of audio encrypted file
	AudioFile string = "audio.bin"

	// END USER CHANGES
)

func Display(_key, _imageFile, _audioFile string, _errMsg *canvas.Text, _window fyne.Window) {
	format.Display(_key, _imageFile, _audioFile, _errMsg, _window)
}
