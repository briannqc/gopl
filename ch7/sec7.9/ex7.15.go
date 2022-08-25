package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main2() {
	input := strings.Join(os.Args[1:], "")
	expr, err := Parse(input)
	if err != nil {
		fmt.Printf("Parsing input failed, err: %v", err)
		os.Exit(1)
	}

	vars := make(map[Var]bool)
	if err := expr.Check(vars); err != nil {
		fmt.Printf("Checking expr failed, err: %v", err)
		os.Exit(1)
	}

	env := Env{}
	if len(vars) > 0 {
		fmt.Println("Expr:", expr)
		fmt.Println("Please enter variable values")

		for v := range vars {
			fmt.Printf("\t%s: ", v)
			var s string
			if _, err := fmt.Scanln(&s); err != nil {
				fmt.Printf("Reading value failed, err: %v", err)
				os.Exit(1)
			}

			value, err := strconv.ParseFloat(s, 64)
			if err != nil {
				fmt.Printf("Parsing value failed, err: %v", err)
				os.Exit(1)
			}
			env[v] = value
		}
	}

	fmt.Println("Value:", expr.Eval(env))
}
