package main

import (
	_ "embed"

	"time"

	"github.com/yajtpg/shrike_go"
)

//go:embed firmware.bin
var firmwareData []byte

func main() {
	time.Sleep(5 * time.Second) // waiting for serial
	shrike_go.Debug = true
	print("[shrikeflash_exp] Starting FPGA flash\n")
	err := shrike_go.Flash(firmwareData)
	if err != nil {
		print("[shrikeflash_exp] Flashing failed\n")
		return
	}
	print("[shrikeflash_exp] Flashing completed successfully\n")
	for {
	}
}
