package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func getAPIKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_KEY := os.Getenv("API_KEY")

	return API_KEY
}

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

var API_KEY string = getAPIKey()

// func GetChallengesList(client *http.Client) []byte {

// 	// List challenges
// 	// List all challenges as well as rewards (xp, stars, cosmetics).
// 	// List of supported languages: en, ar, de, es, es-419, fr, it, ja, ko, pl, pt-BR, ru, tr, zh-CN, zh-Hant

// 	API_KEY := GetAPIKey()

// 	endpoint := "https://fortniteapi.io/v3/challenges"
// 	values := map[string]string{"foo": "baz"}
// 	jsonData, err := json.Marshal(values)

// 	if err != nil {
// 		log.Fatalf("Error Occurred. %+v", err)
// 	}

// 	req, err := http.NewRequest(http.MethodGet, endpoint, bytes.NewBuffer(jsonData))
// 	req.Header.Set("Authorization", API_KEY)
// 	if err != nil {
// 		log.Fatalf("Error Occurred. %+v", err)
// 	}

// 	f := colorjson.NewFormatter()
// 	f.Indent = 4

// 	response, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalf("Error sending request to API endpoint. %+v", err)
// 	}

// 	defer response.Body.Close()

// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatalf("Couldn't parse response body. %+v", err)
// 	}

// 	// s, _ := f.Marshal(body)
// 	// fmt.Println(string(s))

// 	beautifulJsonByte, err := json.MarshalIndent(body, "", "  ")
// 	if err != nil {
// 		panic(err)
// 	}

// 	_ = ioutil.WriteFile("test.json", beautifulJsonByte, 0644)

// 	return body
// 	// return body["bundles"]
// }

// map[string]interface {}

func convert_byte_to_map(input_data []byte) map[string]interface{} {
	var output_data map[string]interface{}

	json.Unmarshal([]byte(input_data), &output_data)
	return output_data

}

func GetChallengesList(client *http.Client) map[string]interface{} {

	// List challenges
	// List all challenges as well as rewards (xp, stars, cosmetics).
	// List of supported languages: en, ar, de, es, es-419, fr, it, ja, ko, pl, pt-BR, ru, tr, zh-CN, zh-Hant
	// en is default!

	endpoint := "https://fortniteapi.io/v3/challenges"

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	req.Header.Set("Authorization", API_KEY)
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	_ = ioutil.WriteFile("challenges_list.json", body, 0644)

	result := convert_byte_to_map(body)

	return result

}

func GetTournamentsList(client *http.Client) map[string]interface{} {

	// Get the list of tournaments
	// List of supported languages: en, ar, de, es, es-419, fr, it, ja, ko, pl, pt-BR, ru, tr
	// en is default!

	endpoint := "https://fortniteapi.io/v1/events/list"

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	req.Header.Set("Authorization", API_KEY)
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	_ = ioutil.WriteFile("tournaments_list.json", body, 0644)

	result := convert_byte_to_map(body)

	return result

}

func main() {
	c := httpClient()

	_ = GetChallengesList(c)

	_ = GetTournamentsList(c)

	fmt.Println("Success!")

	// fmt.Println(challenges_List, )

	// for index, value := range response {
	// 	fmt.Println(index, value)
	// }
	// fmt.Println(reflect.TypeOf(response))
}
