package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/briannqc/gopl/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "conv: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fehrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %s, %s = %s, %s = %s\n", f, f.ToC(), c, c.ToK(), k, k.ToF())
	}
}
