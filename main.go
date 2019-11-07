package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type merchantResponse struct {
	ErrorCode string `xml:"ErrorCode"`
	AuthToken string `xml:"AuthToken"`
}

func main() {
	xmlFile, err := os.Open("merchantResponse.xml")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	var response merchantResponse
	xml.Unmarshal(byteValue, &response)

	fmt.Printf("%#v", response)
}
