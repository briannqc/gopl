package main

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeXML(t *testing.T) {
	tests := []struct {
		name string
		xml  string
		want *Element
	}{
		{
			name: "GIVEN single node xml THEN single node tree",
			xml:  `<a/>`,
			want: &Element{
				Type: xml.Name{
					Local: "a",
				},
				Attr: []xml.Attr{},
			},
		},
		{
			name: "GIVEN single node text xml THEN single node text tree",
			xml:  `<a>A</a>`,
			want: &Element{
				Type: xml.Name{
					Local: "a",
				},
				Attr:     []xml.Attr{},
				Children: append([]Node{}, CharData("A")),
			},
		},
		{
			name: "GIVEN single node with attr xml THEN single node with attr tree",
			xml:  `<a attr="A">V</a>`,
			want: &Element{
				Type: xml.Name{
					Local: "a",
				},
				Attr: []xml.Attr{
					{
						Name: xml.Name{
							Local: "attr",
						},
						Value: "A",
					},
				},
				Children: append([]Node{}, CharData("V")),
			},
		},
		{
			name: "GIVEN 2 level node with attr xml THEN 2 level node with attr tree",
			xml:  `<l0 attr="A"><l1>L1</l1>L0</l0>`,
			want: &Element{
				Type: xml.Name{
					Local: "l0",
				},
				Attr: []xml.Attr{
					{
						Name: xml.Name{
							Local: "attr",
						},
						Value: "A",
					},
				},
				Children: []Node{
					&Element{
						Type: xml.Name{
							Local: "l1",
						},
						Attr:     []xml.Attr{},
						Children: []Node{CharData("L1")},
					},
					CharData("L0"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.xml)
			got, err := DecodeXML(r)
			if err != nil {
				t.Errorf("DecodeXML() returned error = %v", err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
