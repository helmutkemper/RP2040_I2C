// Package blackbox
//
// I2CBus configures a hardware I2C peripheral on a Raspberry Pi Pico (RP2040).
//
// Place at the global scope (outside any Loop). Wire its "bus" output to
// the "i2c" input of any sensor that requires an I2C connection.
package blackbox

import (
	"machine"
	"time"
)

// I2CBus configures an I2C bus. All settings come from the Inspect panel.
// icon:microchip. label:My Device. interactive:rp2040.
type I2CBus struct {
	sda    string `prop:"SDA Pin"   default:"GP4"    options:"GP4,GP6,GP8,GP10,GP12,GP14" pin:"I2C_SDA"  color:"#7c3aed" unit:"Pin"`
	scl    string `prop:"SCL Pin"   default:"GP5"    options:"GP5,GP7,GP9,GP11,GP13,GP15" pin:"I2C_SCL"  color:"#0d9488" unit:"Pin"`
	intPin string `prop:"INT Pin"   default:"GP3"    options:"GP2,GP3,GP22"               pin:"GPIO_INT" color:"#dc2626" unit:"Pin"`
	freq   string `prop:"Frequency" default:"100000" options:"100000,400000,1000000"                                     unit:"bps"`
}

// Init sets up the I2C peripheral and returns a ready bus reference.
//
// executionOrder:1
//
// Returns
//
//	bus: configured I2C bus — wire this to sensors that need I2C.  connection:mandatory.  unit:i2c_bus.
func (s *I2CBus) Init() (bus *machine.I2C) {
	b := machine.I2C0
	b.Configure(machine.I2CConfig{
		Frequency: 400_000,
		SDA:       machine.GP4,
		SCL:       machine.GP5,
	})
	time.Sleep(100 * time.Millisecond)
	return b
}

/*
manualName:wiring-guide.
language:en.
showIn:init.
```markdown
# I2CBus — Wiring Guide

| Signal | Default Pico Pin |
|--------|-----------------|
| SDA    | GP4             |
| SCL    | GP5             |
| VCC    | 3V3 (pin 36)    |
| GND    | GND (pin 38)    |

Wire the **bus** output to the **i2c** input of any sensor block.
```*/
