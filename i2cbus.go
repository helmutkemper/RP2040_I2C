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

// RP2040_I2C configures an I2C bus. All settings come from the Inspect panel.
//
// icon:serial. label:RP2040 I2C. interactive:rp2040.
type RP2040_I2C struct {
	sda  string `prop:"SDA Pin"   default:"GP4"    options:"GP4,GP6,GP8,GP10,GP12,GP14" connection:"I2C_SDA"`
	scl  string `prop:"SCL Pin"   default:"GP5"    options:"GP5,GP7,GP9,GP11,GP13,GP15" connection:"I2C_SCL"`
	freq string `prop:"Frequency" default:"100000" options:"100000,400000,1000000"`
}

// Init sets up the I2C peripheral and returns a ready bus reference.
//
// executionOrder:1. icon:circle-play. label:Init.
//
// Returns
//
//	bus: configured I2C bus — wire this to sensors that need I2C.  connection:mandatory.  unit:i2c_bus.
func (s *RP2040_I2C) Init() (bus *machine.I2C) {
	b := machine.I2C0
	b.Configure(machine.I2CConfig{
		Frequency: 400_000,
		SDA:       machine.GP4,
		SCL:       machine.GP5,
	})
	time.Sleep(100 * time.Millisecond)
	return b
}
