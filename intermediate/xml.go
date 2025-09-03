package basics

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Address Address  `xml:"address"`
	// City    string   `xml:"city"`
	Email string `xml:"email"`
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {

	person := Person{Name: "John", Age: 30, Email: "email@exampleemail.com", Address: Address{City: "Oakland", State: "CA"}}

	// xmlData, err := xml.Marshal(person)
	// if err != nil {
	// 	log.Fatalln("Error Marshalling data into XML:", err)
	// }

	// fmt.Println("XML Data", string(xmlData))

	xmlData1, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatalln("Error Marshalling data into XML:", err)
	}
	fmt.Println("XML Data with Indent:", string(xmlData1))

	// xmlRaw := `<person><name>John</name><age>25</age></person>`
	xmlRaw := `<person><name>John</name><age>25</age><address><city>San Fransisco</city><state>CA</state></address></person>`

	var personxml Person

	err = xml.Unmarshal([]byte(xmlRaw), &personxml)
	if err != nil {
		log.Fatalln("Error Unmarshalling xml", err)
	}
	fmt.Println(personxml)
	fmt.Println("Local", personxml.XMLName.Local)
	fmt.Println("Namespace", personxml.XMLName.Space)

	book := Book{
		ISBN:       "5845-09-534-3434-343-4333",
		Title:      "Go Bootcamp",
		Author:     "Alfred",
		Pseudo:     "Pseudo",
		PseudoAttr: "Pseudo Attribute",
	}

	xmlDatAttr, err := xml.MarshalIndent(book, "a", "  ")
	if err != nil {
		log.Fatalln("Error marshalling data:", err)
	}

	fmt.Println(string(xmlDatAttr))

}

type Book struct {
	XMLName    xml.Name `xml:"book"`
	ISBN       string   `xml:"isbn,attr"`
	Title      string   `xml:"title,attr"`
	Author     string   `xml:"author,attr"`
	Pseudo     string   `xml:"pseudo"`
	PseudoAttr string   `xml:"pseudoattr,attr"`
}

// <book isbn="jngfrrcrcrcedcedc7ytrf" color="blue">
