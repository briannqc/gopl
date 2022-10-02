package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func main() {
	outFormat := flag.String("f", "jpeg", "Output format")
	flag.Parse()

	fn, ok := map[string]encodeFunc{
		"jpeg": encodeJPEG,
		"jpg":  encodeJPEG,
		"png":  png.Encode,
		"gif":  encodeGIF,
	}[*outFormat]

	if !ok {
		_, _ = fmt.Fprintf(os.Stderr, "Format: %v is not supported", *outFormat)
		os.Exit(1)
	}

	img, kind, err := image.Decode(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Decoding input failed, err: %v", err)
		os.Exit(1)
	}

	_, _ = fmt.Fprintf(os.Stderr, "Converting from %v to %v\n", kind, *outFormat)
	if err := fn(os.Stdout, img); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Converting failed, err: %v\n", err)
		os.Exit(1)
	}
	_, _ = fmt.Fprintln(os.Stderr, "Converting succeeded")
}

type encodeFunc func(out io.Writer, img image.Image) error

func encodeJPEG(out io.Writer, img image.Image) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func encodeGIF(out io.Writer, img image.Image) error {
	return gif.Encode(out, img, &gif.Options{})
}
