package tempconv

import "fmt"

type Celsius float64
type Fehrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

func (c Celsius) ToF() Fehrenheit {
	return Fehrenheit(c*9/5 + 32)
}

func (c Celsius) ToK() Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

func (f Fehrenheit) String() string {
	return fmt.Sprintf("%.2f°F", f)
}

func (f Fehrenheit) ToC() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (f Fehrenheit) ToK() Kelvin {
	return f.ToC().ToK()
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%.2fK", k)
}

func (k Kelvin) ToC() Celsius {
	return Celsius(k) + AbsoluteZeroC
}

func (k Kelvin) ToF() Fehrenheit {
	return k.ToC().ToF()
}
