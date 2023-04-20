package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type RequestBody struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		apiUrl := "https://jsonplaceholder.typicode.com/posts"

		data := RequestBody{
			Water: rand.Intn(16),
			Wind:  rand.Intn(16),
		}

		bs, err := json.Marshal(data)

		if err != nil {
			log.Panicf("error while converting struct to json => %s \n", err.Error())
		}

		request, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(bs))

		if err != nil {
			log.Panicf("error while defining the request instance => %s \n", err.Error())
		}

		request.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		response, err := client.Do(request)

		defer response.Body.Close()

		if err != nil {
			log.Panicf("error while sending the api request => %s \n", err.Error())
		}

		responseBody, err := ioutil.ReadAll(response.Body)

		fmt.Println(string(responseBody))

		if data.Water >= 6 && data.Water <= 8 {
			fmt.Println("status water : siaga")
		} else if data.Water < 5 {
			fmt.Println("status water: aman")
		}

		time.Sleep(time.Second * 2)
	}
}
