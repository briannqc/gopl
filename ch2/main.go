package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/briannqc/gopl/ch2/lenconv"
	"github.com/briannqc/gopl/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "conv: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fehrenheit(v)
		c := tempconv.Celsius(v)
		k := tempconv.Kelvin(v)
		fmt.Printf("%s = %s, %s = %s, %s = %s\n", f, f.ToC(), c, c.ToK(), k, k.ToF())

		cm := lenconv.Centimeter(v)
		in := lenconv.Inch(v)
		fmt.Printf("%s = %s, %s = %s\n", cm, cm.ToInch(), in, in.ToCentimeter())
	}
}
