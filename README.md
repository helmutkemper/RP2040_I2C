# RP2040 I2C

Este código configura a I2C da placa Raspberry Pi RP2040.

## Propriedades

```go
type APDS9960 struct {
	sda  string `prop:"SDA Pin"   default:"GP4"    options:"GP4,GP6,GP8,GP10,GP12,GP14" pin:"I2C_SDA"  unit:"Pin"`
	scl  string `prop:"SCL Pin"   default:"GP5"    options:"GP5,GP7,GP9,GP11,GP13,GP15" pin:"I2C_SCL"  unit:"Pin"`
	intPin string `prop:"INT Pin" default:"GP3"    options:"GP2,GP3,GP22"               pin:"GPIO_INT" unit:"Pin"`
	freq string `prop:"Frequency" default:"400000" options:"100000,400000,1000000"                     unit:"bps"`
}
```

**Painel de configurações:**
```
  +----------------------------+
  |  SDA Pin   [   GP4 ▼] Pin  |
  |  SCL Pin   [   GP5 ▼] Pin  |
  |  INT Pin   [   GP3 ▼] Pin  |
  |  Frequency [100000 ▼] bps  |
  +----------------------------+
```

## Placa RP2040

![](rp2040.svg)
