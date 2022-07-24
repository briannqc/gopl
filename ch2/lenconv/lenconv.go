package lenconv

import "fmt"

type Centimeter float64

type Inch float64

func (c Centimeter) ToInch() Inch {
	return Inch(c / 2.54)
}

func (c Centimeter) String() string {
	return fmt.Sprintf("%.2f cm", c)
}

func (i Inch) ToCentimeter() Centimeter {
	return Centimeter(i * 2.54)
}

func (i Inch) String() string {
	return fmt.Sprintf("%.2f inch", i)
}
