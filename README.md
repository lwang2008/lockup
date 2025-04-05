# Lockup (Archived 2021)

If you create NFTs (Non-fungible tokens), Lockup is a way to allow only the owner of the file to be able to view the image.

## Note About Repository History

This is a fresh repository created in 2025. The original repository contained sensitive information and was therefore archived. The code was originally developed in 2021, and the copyright notices reflect the original development date. There is a lot of deprecated code which may render the program unusable. I am currently in the process of making these improvements.

## Using Lockup
### Locking up a file

NOTE: This needs to be updated, the whole process has been greatly simplified.

DISCLAIMER: Lockup's point of failure is with an entity who holds the unencrypted, orginal, copy of the file. If this file is published, using Lockup would be pointless, so be sure this image is stored securely offline if it holds value.

Prerequisites: 
- Golang is installed on your system. If you haven't done so already, you can click this [link](https://golang.org/doc/install) to learn how.

Process:
1. Open a new terminal window and run this command `curl https:// <LINK> -O lockup.zip`. Then run `unzip -D lockup` and finally `cd lockup` (Check the releases page to see which link you should use)
2. We now need to encrypt the image file.
    1. Put the image in the lockup folder and then in the `encrypt` folder. 
    2. Click/run the `encrypt` file in that folder. Complete the steps required.
    3. Be sure to copy the file name perfectly and jot down your encryption key for use later.
    4. Suggestion: Give the encrypted file a `.bin` suffix to indicate an encrypted binary.
    5. Move the encrypted file back to the `lockup` folder.
3. Finally we will create the application binary that will decrypt and display the image.
    1. Navigate to the `params.go` file in the `lockup/params` folder your preferred text editor
    2. If using a file format that is not an image(i.e. audio or music), comment that line (add a `//` to the beginning of the line) and uncomment the line of your preferred format (remove the relevant `//`.
    3. Edit the lines of code beneath `// START USER CHANGES` and `// END USER CHANGES`
        1. `tokenID` is the ID of the specific NFT of the smart contract you will use.
        2. `contractAddress` is the smart contract that the NFT is minted on.
        Here are the addresses of some common platforms.
            | Platform | Contract |
            |:-:|:-:|
            |Foundation|[`0x3B3ee1931Dc30C1957379FAc9aba94D1C48a5405`](https://etherscan.io/address/0x3B3ee1931Dc30C1957379FAc9aba94D1C48a5405)|
            |SuperRare| [`0xb932a70a57673d89f4acffbe830e8ed7f75fb9e0`](https://etherscan.io/address/0xb932a70a57673d89f4acffbe830e8ed7f75fb9e0)|
            |Async.art| [`0xb6dae651468e9593e4581705a09c10a76ac1e0c8`](https://etherscan.io/address/0xb6dae651468e9593e4581705a09c10a76ac1e0c8)|
            |Rarible Collection| [`0x60f80121c31a0d46b5279700f9df786054aa5ee5`](https://etherscan.io/address/0x60f80121c31a0d46b5279700f9df786054aa5ee5)|
        3. `key` is the encryption key that you used to encrypt the file during step 2.
        4. `rpcUrl` is the URL of an Ethereum node's JSON RPC, this is to find the owner of a specific NFT. You can use services such as [Infura](https://infura.io) or [Alchemy](https://www.alchemy.com/). You can also run your own public node, just make sure it is secured with SSL and the server is publicly accessible.
        5. `nftTitle` is the title of your NFT. It will be displayed as the title of the window.
        6. `fileName` is just the encrypted file of your NFT
    3. After editing the file, save it, return to the console and run `go build`
    4. Finally, make a new folder and put the `lockup` file (not `lockup.go`) and the `.bin` file.
4. Test your by following the instructions in the next section.
5. Finally, if you wish to distribute this file, zip the folder and upload it to whatever service you choose such as [IPFS](https://ipfs.io) or [Swarm](https://swarm.ethereum.org)

### Viewing/"Unlocking" a file
1. Click the `lockup` file.
2. You will need to sign a digital signature. A terminal window should open upon clicking the file, it will contain the message you will need to sign.
3. Sign the message. You can use any web3 wallet or signing tool that supports the `personal_sign` method. If you would like to sign it yourself use the [`personal_sign`](https://geth.ethereum.org/docs/rpc/ns-personal#personal_sign) instead of `eth_sign` method, this is to avoid malcious programs from getting you to sign transactions.
4. Enter the in the gui window. Make sure the data has a `0x` prefix.
5. If successful, your image should be shown.

## Architecture
Lockup is meant to restrict access to the original NFT images in an almost completely decentralized manner. The system first calls an ERC721 contract to find the owner of an NFT. Then it uses a digital signature to verify the token owner. If the signature is valid, the program will decrypt the image and store it in a base64 format io stream (not a file). The io stream is read by the Fyne GUI API to display it in the window.

## Goals
- [ ] Expand encryption to video and audio
