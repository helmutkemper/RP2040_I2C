# Configuration

```go
type RP2040_I2C struct {
	sda    string `prop:"SDA Pin"   default:"GP4"    options:"GP4,GP6,GP8,GP10,GP12,GP14" pin:"I2C_SDA"  color:"#7c3aed" unit:"Pin"`
	scl    string `prop:"SCL Pin"   default:"GP5"    options:"GP5,GP7,GP9,GP11,GP13,GP15" pin:"I2C_SCL"  color:"#0d9488" unit:"Pin"`
	intPin string `prop:"INT Pin"   default:"GP3"    options:"GP2,GP3,GP22"               pin:"GPIO_INT" color:"#dc2626" unit:"Pin"`
	freq   string `prop:"Frequency" default:"100000" options:"100000,400000,1000000"                                     unit:"bps"`
}
```

**Settings panel:**
```
  +----------------------------+
  |  SDA Pin   [   GP4 ▼] Pin  |
  |  SCL Pin   [   GP5 ▼] Pin  |
  |  INT Pin   [   GP3 ▼] Pin  |
  |  Frequency [100000 ▼] bps  |
  +----------------------------+
```

## Board RP2040

![](rp2040.svg)

