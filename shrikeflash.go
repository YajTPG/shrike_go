package shrike_go

// Ported from: https://github.com/vicharak-in/shrike_arduino/blob/main/src/Shrike.cpp

import (
	"machine"
	"time"
)

var (
	SPI  = machine.SPI0
	_PWR = machine.FPGA_PWR
	_EN  = machine.FPGA_EN

	_SS   = machine.GP1
	_SCK  = machine.GP2
	_MOSI = machine.GP3
	_MISO = machine.GP0
)

var Debug = false

func init() {
	SPI.Configure(machine.SPIConfig{
		SDO:       _MOSI,
		SDI:       _MISO,
		SCK:       _SCK,
		Frequency: 1600000,
		Mode:      0b00,
	})
	_PWR.Configure(machine.PinConfig{Mode: machine.PinOutput})
	_EN.Configure(machine.PinConfig{Mode: machine.PinOutput})
	_SS.Configure(machine.PinConfig{Mode: machine.PinOutput})

	_SS.High()
	_EN.Low()
	_PWR.Low()
}

func Flash(data []byte) error {
	log("Starting FPGA flash")

	Reset()

	time.Sleep(100 * time.Millisecond)
	_SS.Low()
	_EN.Low()
	_PWR.Low()
	time.Sleep(100 * time.Millisecond)
	_PWR.High()
	_EN.High()
	time.Sleep(100 * time.Millisecond)

	_SS.High()
	time.Sleep(2000 * time.Microsecond)
	_SS.Low()

	err := SPI.Tx(data, nil)
	if err != nil {
		print("Error writing to SPI")
		return (err)
	}

	log("Flash successful")
	_SS.High()
	time.Sleep(100 * time.Microsecond)
	return nil
}

func Reset() {
	_PWR.Low()
	_EN.Low()
	print("FPGA Reset")
}

func log(msg string) {
	if Debug {
		print("[shrikeflash] " + msg + "\n")
	}
}
