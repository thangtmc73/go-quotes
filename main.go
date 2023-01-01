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

const quoteURL = "https://gist.githubusercontent.com/thangtmc73/aefb6536e4869f89e914e7c70b0269c3/raw/94c565c0932c4a25180fcba6e3bfdfed6c6b07cf/go_quotes.json"

func main() {
	quotes, err := getQuotesFromLocalJSON("default_quotes.json")
	if err != nil {
		fmt.Println("Could not get quotes from local JSON", err)
	}
	quoteIndex := rand.Intn(len(quotes))
	quote := quotes[quoteIndex]

	fmt.Println("+-------------+")
	fmt.Println("| ", quote.Content, " |")
	fmt.Println("|-", quote.Author, " |")
	fmt.Println("+-------------+")
}

func getQuotesFromLocalJSON(filePath string) ([]Quote, error) {
	var quotes []Quote
	jsonFile, openErr := os.Open(filePath)
	defer func(jsonFile *os.File) {
		closeErr := jsonFile.Close()
		if closeErr != nil {
			fmt.Println("Could not close file", closeErr)
		}
	}(jsonFile)
	if openErr != nil {
		return quotes, openErr
	}
	byteValue, _ := io.ReadAll(jsonFile)
	if parseErr := json.Unmarshal(byteValue, &quotes); parseErr != nil {
		return quotes, parseErr
	}
	return quotes, nil
}
