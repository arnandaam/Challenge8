package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"math/rand"
// 	"net/http"
// 	"strconv"
// 	"time"
// )

// type Data struct {
// 	Water string `json:"valuewater"`
// 	Wind  string `json:"valuewind"`
// }

// type Response struct {
// 	ID     int    `json:"id"`
// 	UserID int    `json:"userId"`
// 	Title  string `json:"title"`
// 	Body   string `json:"body"`
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	for {
// 		data := Data{
// 			Water: strconv.Itoa(rand.Intn(100) + 1),
// 			Wind:  strconv.Itoa(rand.Intn(100) + 1),
// 		}
// 		jsonValue, _ := json.Marshal(data)

// 		fmt.Println("Request JSON: ", string(jsonValue))

// 		response, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonValue))
// 		if err != nil {
// 			fmt.Println("Error posting data:", err)
// 			return
// 		}
// 		defer response.Body.Close()

// 		responseBody, _ := ioutil.ReadAll(response.Body)
// 		fmt.Println("Response JSON: ", string(responseBody))

// 		var result Response
// 		json.Unmarshal(responseBody, &result)

// 		fmt.Println("Water Status: ", getStatusWater(data.Water))
// 		fmt.Println("Wind Status: ", getStatusWind(data.Wind))
// 		fmt.Println("Response ID: ", result.ID)
// 		fmt.Println("-------------")
// 		time.Sleep(15 * time.Second)
// 	}
// }

// func getStatusWater(water string) string {
// 	value, _ := strconv.Atoi(water)
// 	if value < 5 {
// 		return "AMAN"
// 	} else if value >= 5 && value <= 8 {
// 		return "SIAGA"
// 	} else {
// 		return "BAHAYA"
// 	}
// }

// func getStatusWind(wind string) string {
// 	value, _ := strconv.Atoi(wind)
// 	if value < 6 {
// 		return "AMAN"
// 	} else if value >= 6 && value <= 15 {
// 		return "SIAGA"
// 	} else {
// 		return "BAHAYA"
// 	}
// }
