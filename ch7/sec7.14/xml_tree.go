package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func DecodeXML(r io.Reader) (*Element, error) {
	decoder := xml.NewDecoder(r)
	var root *Element
	var stack []*Element
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			return nil, err
		}

		switch token := token.(type) {
		case xml.StartElement:
			child := &Element{
				Type: token.Name,
				Attr: token.Attr,
			}
			if root == nil {
				root = child
			} else {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, child)
			}
			stack = append(stack, child)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			parent := stack[len(stack)-1]
			child := CharData(token)
			parent.Children = append(parent.Children, child)
		}
	}
	return root, nil
}
