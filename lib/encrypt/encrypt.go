// Copyright 2021, Lucas Wang
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"os"
)

func Encrypt(_key, _pname string) ([]byte, error) {

	for len(_key) < 32 {
		_key += "0"
	}

	plaintext, err := os.ReadFile(_pname)

	if err != nil {
		return nil, err

	}

	block, err := aes.NewCipher([]byte(_key))

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	return ciphertext, nil

}
