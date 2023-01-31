package main

import (
	"machine"

	//	"time"

	"tinygo.org/x/drivers/lora"
	"tinygo.org/x/drivers/sx127x"
)

const (
	rxTimeoutMs = 2000
	txTimeoutMs = 2000
)

var (
	loraRadio *sx127x.Device

	// LoRa Featherwing module is connected to PyBadge:
	SX127X_PIN_RST  = machine.D11
	SX127X_PIN_CS   = machine.D10
	SX127X_PIN_DIO0 = machine.D6
	SX127X_PIN_DIO1 = machine.D9
	SX127X_SPI      = machine.SPI0
)

func startLora() {
	SX127X_PIN_RST.Configure(machine.PinConfig{Mode: machine.PinOutput})
	SX127X_SPI.Configure(machine.SPIConfig{Frequency: 500000, Mode: 0})

	loraRadio = sx127x.New(SX127X_SPI, SX127X_PIN_RST)
	loraRadio.SetRadioController(sx127x.NewRadioControl(SX127X_PIN_CS, SX127X_PIN_DIO0, SX127X_PIN_DIO1))

	loraRadio.Reset()
	state := loraRadio.DetectDevice()
	if !state {
		panic("main: sx127x NOT FOUND !!!")
	} else {
		println("main: sx127x found")
	}

	// Prepare for Lora Operation
	loraConf := lora.Config{
		Freq:           lora.MHz_868_1,
		Bw:             lora.Bandwidth_125_0,
		Sf:             lora.SpreadingFactor9,
		Cr:             lora.CodingRate4_7,
		HeaderType:     lora.HeaderExplicit,
		Preamble:       12,
		Iq:             lora.IQStandard,
		Crc:            lora.CRCOn,
		SyncWord:       lora.SyncPrivate,
		LoraTxPowerDBm: 20,
	}

	loraRadio.LoraConfig(loraConf)
}

func loraRX() {
	for {
		buf, err := loraRadio.Rx(rxTimeoutMs)
		switch {
		case err != nil:
			println("RX Error: ", err)
			showError(err)
		case buf == nil || len(buf) == 0:
			// empty buffer, do nothing
		case buf[0] == '@' && buf[len(buf)-1] == '!':
			showMessage(buf[:len(buf)-1])
		default:
			println("Unknown packet received: len=", len(buf))
		}
	}
}

func loraTX(msg []byte) error {
	err := loraRadio.Tx(msg, txTimeoutMs)
	if err != nil {
		return err
	}

	return nil
}
