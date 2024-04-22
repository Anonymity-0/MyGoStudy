package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	// XMLName xml.Name `xml:"plant"` // 如果这里是 plant，那么输出的xml就是 <plant>...</plant>
	XMLName xml.Name `xml:"plane"`
	// id,attr 表示这个字段是一个 id 属性。
	Id    int      `xml:"id,attr"`
	Name  string   `xml:"name"`
	Orgin []string `xml:"orgin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, orgin=%v", p.Id, p.Name, p.Orgin)
}
func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Orgin = []string{"Ethiopia", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, "", "  ")
	fmt.Println(string(out))

	fmt.Println(xml.Header + string(out))

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Orgin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = append(nesting.Plants, coffee)
	nesting.Plants = append(nesting.Plants, tomato)

	out, _ = xml.MarshalIndent(nesting, "", "  ")
	fmt.Println(string(out))

}
