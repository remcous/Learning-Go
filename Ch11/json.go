package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	toFile := Person{
		Name: "Fred",
		Age:  40,
	}

	tmpFile, err := ioutil.TempFile(os.TempDir(), "sample-")
	if err != nil {
		panic(err)
	}

	defer os.Remove(tmpFile.Name())

	err = json.NewEncoder(tmpFile).Encode(toFile)
	if err != nil {
		panic(err)
	}

	err = tmpFile.Close()
	if err != nil {
		panic(err)
	}

	tmpFile2, err := os.Open(tmpFile.Name())
	if err != nil {
		panic(err)
	}

	var fromFile Person
	err = json.NewDecoder(tmpFile2).Decode(&fromFile)
	if err != nil {
		panic(err)
	}

	err = tmpFile2.Close()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", fromFile)

	const data = `
		{"name": "Fred", "age": 40}
		{"name": "Mary", "age": 21}
		{"name": "Pat", "age": 30}
	`

	var t struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	dec := json.NewDecoder(strings.NewReader(data))
	var b bytes.Buffer
	enc := json.NewEncoder(&b)

	for dec.More() {
		err := dec.Decode(&t)
		if err != nil {
			panic(err)
		}

		fmt.Println(t)

		err = enc.Encode(t)
		if err != nil {
			panic(err)
		}
	}
}
