package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Data struct {
	Water string `json:"valuewater"`
	Wind  string `json:"valuewind"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		data := Data{
			Water: strconv.Itoa(rand.Intn(100)),
			Wind:  strconv.Itoa(rand.Intn(100)),
		}
		jsonValue, _ := json.Marshal(data)

		response, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			fmt.Println("Error posting data:", err)
			return
		}
		defer response.Body.Close()

		responseBody, _ := ioutil.ReadAll(response.Body)
		fmt.Println("Response JSON: ", string(responseBody))

		var result map[string]interface{}
		json.Unmarshal(responseBody, &result)

		fmt.Println("Water Status: ", getStatusWater(data.Water))
		fmt.Println("Wind Status: ", getStatusWind(data.Wind))
		time.Sleep(15 * time.Second)
	}
}

func getStatusWater(water string) string {
	value, _ := strconv.Atoi(water)
	if value < 5 {
		return "AMAN"
	} else if value >= 5 && value <= 8 {
		return "SIAGA"
	} else {
		return "BAHAYA"
	}
}

func getStatusWind(wind string) string {
	value, _ := strconv.Atoi(wind)
	if value < 6 {
		return "AMAN"
	} else if value >= 6 && value <= 15 {
		return "SIAGA"
	} else {
		return "BAHAYA"
	}
}
