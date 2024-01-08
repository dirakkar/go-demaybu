// Copyright (c) 2023-2024 Dirakkar
// SPDX-License-Identifier: MIT

package crypto

import (
	"github.com/zeebo/blake3"

	"github.com/dirakkar/go-demaybu/dapp"
)

const SchemaED25519 = 0x1

func NewAddress(schema byte, publicKey []byte) dapp.Address {
	tmp := []byte{schema}
	tmp = append(tmp, publicKey...)

	return dapp.Address(blake3.Sum256(tmp))
}
