package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Quote []QuoteElement

type QuoteElement struct {
	Q string `json:"q"`
	A string `json:"a"`
	H string `json:"h"`
}

func main() {
	for {
		response, err := http.Get("https://zenquotes.io/api/random")
		if err != nil {
			panic(err)
		}

		data, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		var quotes Quote
		err = json.Unmarshal(data, &quotes)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s ~ by %s\n", quotes[0].Q, quotes[0].A)
		time.Sleep(time.Second * 5)
	}
}
