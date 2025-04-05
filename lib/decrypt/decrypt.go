// Copyright 2021, Lucas Wang
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"os"
)

func Decrypt(_key, _ename string) (string, error) {

	for len(_key) < 32 {
		_key += "0"
	}

	ciphertext, err := os.ReadFile(_ename)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(_key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), err
}
