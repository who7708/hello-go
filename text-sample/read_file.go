package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

func getAttributeValue(attr []xml.Attr, name string) string {
	for _, a := range attr {
		if a.Name.Local == name {
			return a.Value
		}
	}
	return ""
}

func readAllFile() {
	// ioutil.ReadFile("read_file_test.iml")
	content, err := os.ReadFile("read_file_test.iml")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	decoder := xml.NewDecoder(bytes.NewBuffer(content))

	var t xml.Token
	var component bool
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			// fmt.Printf("name: %v\n", name)
			if component {
				if name == "orderEntry" {
					// fmt.Printf("name: %v\n", name)
					s := getAttributeValue(token.Attr, "name")
					if strings.TrimSpace(s) == "" {
						continue
					}
					s = strings.Replace(s, "Maven: ", "", -1)
					fmt.Printf("getAttributeValue(token.Attr, \"name\"): %v\n", s)
				}
			} else {
				if name == "component" {
					component = true
				}
			}
		case xml.EndElement:
			if component {
				if token.Name.Local == "component" {
					component = false
				}
			}
		}
	}

}
