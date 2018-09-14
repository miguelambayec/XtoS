package main

import (
	"encoding/xml"
	"fmt"
	"xtos/test/struct"
)

func main() {

	sample1 := model.Bookstore{}

	xmlByte, _ := xml.MarshalIndent(sample1, "  ", "    ")

	fmt.Println(string(xmlByte))
}
