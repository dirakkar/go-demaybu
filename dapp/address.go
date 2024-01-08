// Copyright (c) 2023-2024 Dirakkar
// SPDX-License-Identifier: MIT

package dapp

import "encoding/hex"

const AddressSize = 32

type Address [AddressSize]byte

func (a Address) String() string {
	return "0x" + hex.EncodeToString(a[:])
}
