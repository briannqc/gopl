package anonymousfunc

import (
	"fmt"
	"testing"
)

func squares() func() int {
	var x int
	fmt.Println(&x)
	return func() int {
		fmt.Println(&x)
		x++
		return x * x
	}
}

/**
=== RUN   TestSquare
0xc000016298
0xc000016298
1
0xc000016298
4
0xc000016298
9
0xc000016298
16
--- PASS: TestSquare (0.00s)
*/
func TestSquare(t *testing.T) {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
