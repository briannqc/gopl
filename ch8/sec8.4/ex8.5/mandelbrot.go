// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"sync"
)

func render(w io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	_ = png.Encode(w, img) // NOTE: ignoring errors
}

func renderInParallel1(w io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	type Point struct {
		x, y  int
		color color.Color
	}

	ch := make(chan Point, height)
	var wg sync.WaitGroup
	for py := 0; py < height; py++ {
		wg.Add(1)

		go func(py int) {
			defer wg.Done()

			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)

				ch <- Point{
					x:     px,
					y:     py,
					color: mandelbrot(z),
				}
			}
		}(py)

	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for p := range ch {
		img.Set(p.x, p.y, p.color)
	}

	_ = png.Encode(w, img) // NOTE: ignoring errors
}

func renderInParallel2(w io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	type Point struct {
		x, y  int
		color color.Color
	}

	ch := make(chan Point, height*width)
	var wg sync.WaitGroup
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			wg.Add(1)

			go func(px, py int) {
				defer wg.Done()

				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)

				ch <- Point{
					x:     px,
					y:     py,
					color: mandelbrot(z),
				}
			}(px, py)
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for p := range ch {
		img.Set(p.x, p.y, p.color)
	}

	_ = png.Encode(w, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}
	return color.Black
}
