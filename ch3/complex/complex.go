package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/big"
	"math/cmplx"
	"os"
)

func main() {
	renderBigRat(os.Stdout)
}

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func renderComplex128(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrotComplex128(z))
		}
	}
	png.Encode(w, img)
}

func mandelbrotComplex128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func renderComplex64(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			var z complex64 = complex(x, y)
			img.Set(px, py, mandelbrotComplex64(z))
		}
	}
	png.Encode(w, img)
}

func mandelbrotComplex64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func renderBigFloat(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			img.Set(px, py, mandelbrotBigFloat(big.NewFloat(x), big.NewFloat(y)))
		}
	}
	png.Encode(w, img)
}

func mandelbrotBigFloat(re, im *big.Float) color.Color {
	const iterations = 200
	const contrast = 15

	vRe, vIm := &big.Float{}, &big.Float{}

	two := big.NewFloat(2)
	four := big.NewFloat(4)
	for n := uint8(0); n < iterations; n++ {
		// (r + i)^2 = r^2 + 2ri + i^2
		vR2, vI2 := &big.Float{}, &big.Float{}
		vR2.Mul(vRe, vRe).Sub(vR2, (&big.Float{}).Mul(vIm, vIm)).Add(vR2, re)
		vI2.Mul(vRe, vIm).Mul(vI2, two).Add(vI2, im)
		vRe, vIm = vR2, vI2

		absSquare := &big.Float{}
		absSquare.Mul(vRe, vRe).Add(absSquare, (&big.Float{}).Mul(vIm, vIm))
		if absSquare.Cmp(four) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func renderBigRat(w io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrotBigRat(z))
		}
	}
	png.Encode(w, img)
}

func mandelbrotBigRat(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	re := (&big.Rat{}).SetFloat64(real(z))
	im := (&big.Rat{}).SetFloat64(imag(z))
	vRe, vIm := &big.Rat{}, &big.Rat{}

	two := big.NewRat(2, 1)
	four := (&big.Rat{}).SetFloat64(4)
	for n := uint8(0); n < iterations; n++ {
		// (r + i)^2 = r^2 + 2ri + i^2
		vR2, vI2 := &big.Rat{}, &big.Rat{}
		vR2.Mul(vRe, vRe).Sub(vR2, (&big.Rat{}).Mul(vIm, vIm)).Add(vR2, re)
		vI2.Mul(vRe, vIm).Mul(vI2, two).Add(vI2, im)
		vRe, vIm = vR2, vI2

		absSquare := &big.Rat{}
		absSquare.Mul(vRe, vRe).Add(absSquare, (&big.Rat{}).Mul(vIm, vIm))
		if absSquare.Cmp(four) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
