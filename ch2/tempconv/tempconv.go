package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

func (c Celsius) ToF() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (c Celsius) ToK() Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", f)
}

func (f Fahrenheit) ToC() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (f Fahrenheit) ToK() Kelvin {
	return f.ToC().ToK()
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%.2fK", k)
}

func (k Kelvin) ToC() Celsius {
	return Celsius(k) + AbsoluteZeroC
}

func (k Kelvin) ToF() Fahrenheit {
	return k.ToC().ToF()
}
