# XLSX Stream 

## Motivation

There are other Excel parser the issue is that they try to load the full file into memmory and when the file is big with millions of records that is not an option, so I make this library that streams line by line using low memmory and CPU.

## Limitations

This library assumes that a given sheet in an xlsx format spreadsheet has nothing but a table in it. The equivelant of a CSV but in a spreadsheet. This library can can only read line by line for processing or loading into databases.

## Usage

```go
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"strings"

	xlsx "github.com/jacostaperu/xlsx-stream"
)

func main() {
	f, err := zip.OpenReader("xlsx_test.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := xlsx.NewReader(f)
	r.Worksheet = "sheet2"

	for {
		row, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(strings.Join(row, ","))
	}
}
```
