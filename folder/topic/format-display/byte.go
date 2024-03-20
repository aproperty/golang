package main

import (
	"fmt"
	"strconv"
)

func BytesToReadableBinaryString(bs []byte) string {
	hex := ""
	for i := 0; i < len(bs); i++ {
		hex = hex + " 0x" + strconv.FormatUint(uint64(bs[i]), 16) + " "
	}
	str := "bytes{" + hex + "}"
	return str
}

func PrintByteArrayToHex(data []byte) {
	for i := 0; i < len(data); i++ {
		fmt.Printf("%02X ", data[i])
	}
	fmt.Println("")
}
