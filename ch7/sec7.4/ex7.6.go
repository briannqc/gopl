package main

import (
	"flag"
	"fmt"
)

// Celsius represents temperature in Celsius.
// It implements flag.Value interface.
type Celsius struct {
	value float64
}

// String returns temperature value with °C.
// That answer exercise 7.7:  Explain why the
// help message contains °C when the default
// value of 20.0 does not.
func (c *Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c.value)
}

func (c *Celsius) Set(v string) error {
	var unit string
	var value float64
	_, _ = fmt.Sscanf(v, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		c.value = value
		return nil
	case "F", "°F":
		c.value = (value - 32) * 5 / 9
		return nil
	case "K":
		c.value = value - 273.15
		return nil
	}
	return fmt.Errorf("invalid temperature %q", v)
}

// CelsiusFlag is similar to a bunch of functions in flag package, e.g. flag.String().
// It defines a Celsius flag with the specific name, default value, and usage, and
// returns the address of the flag variable. The flag argument must have a quantity
// and a unit, e.g., "100C".
// Supported units are C/°C, F/°F and K.
func CelsiusFlag(name string, defaultValue Celsius, usage string) *Celsius {
	c := defaultValue
	flag.CommandLine.Var(&c, name, usage)
	return &c
}

func main() {
	temp := CelsiusFlag("temp", Celsius{value: 20.0}, "Temperature")
	flag.Parse()

	fmt.Println(temp)
}
