package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
)

type Quote struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, openErr := os.Open("default_quotes.json")
	defer func(jsonFile *os.File) {
		closeErr := jsonFile.Close()
		if closeErr != nil {
			fmt.Println("Close file error", closeErr)
		}
	}(jsonFile)
	if openErr != nil {
		fmt.Println("Open file error", openErr)
	}
	var quotes []Quote
	byteValue, _ := io.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &quotes); err != nil {
		panic(err)
	}

	quoteIndex := rand.Intn(len(quotes))
	quote := quotes[quoteIndex]

	fmt.Println("+-------------+")
	fmt.Println("| ", quote.Content, " |")
	fmt.Println("|-", quote.Author, " |")
	fmt.Println("+-------------+")
}
