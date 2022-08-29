package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	selectBy := flag.String("select-by", "name", "name/attributes")
	flag.Parse()

	dec := xml.NewDecoder(os.Stdin)
	if *selectBy == "attributes" {
		selectByAttributes(dec)
	} else {
		selectByName(dec)
	}
}

func selectByAttributes(dec *xml.Decoder) {
	args := flag.Args()
	attrs := map[string]string{}
	for _, arg := range args {
		kv := strings.Split(arg, "=")
		k, v := kv[0], kv[1]
		attrs[k] = v
	}

	var stack [][]xml.Attr
	for {
		token, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch token := token.(type) {
		case xml.StartElement:
			stack = append(stack, token.Attr)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if len(stack) > 0 {
				last := stack[len(stack)-1]
				if containsAllAttributes(last, attrs) {
					fmt.Printf("%s: %s\n", last, token)
				}
			}
		}
	}
}

func containsAllAttributes(all []xml.Attr, sub map[string]string) bool {
	allAttrs := map[string]string{}
	for _, a := range all {
		allAttrs[a.Name.Local] = a.Value
	}

	for name, value := range sub {
		if allAttrs[name] != value {
			return false
		}
	}
	return true
}

func selectByName(dec *xml.Decoder) {
	var stack []string
	for {
		token, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch token := token.(type) {
		case xml.StartElement:
			stack = append(stack, token.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, flag.Args()) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), token)
			}
		}
	}
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
