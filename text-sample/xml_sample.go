package main

import (
	"encoding/xml"
	"fmt"
)

// 人物档案
type Person struct {
	Name     string `xml:"name,attr"`
	Location string
	Age      int `xml:"age"`
}

func convertPersonByXml() {
	p := &Person{
		Name:     "zhangsa",
		Location: "shanghai",
		Age:      33,
	}

	var data []byte
	var err error

	// 序列化成 xml
	// if data, err := xml.Marshal(p); err != nil {
	// MarshalIndent 第2个参数前缀，第3个参数缩进
	if data, err = xml.MarshalIndent(p, "", "  "); err != nil {
		fmt.Println(err)
		return
	} else {
		// <Person name="zhangsa">
		//   <Location>shanghai</Location>
		//   <age>33</age>
		// </Person>
		fmt.Println(string(data))
	}

	p2 := new(Person)

	if err = xml.Unmarshal(data, &p2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p2)
}
