package main

import "fmt"
import "net/http"
import "io/ioutil"
import "log"
import "os"
import "github.com/joho/godotenv"

func main() {

	// Get vars from .env file
	err := godotenv.Load()
	if err != nil {
    	log.Fatal("Error loading .env file")
  	}

	api_key := os.Getenv("API_KEY")
	location_id := os.Getenv("LOCATION_ID")
	units := os.Getenv("UNITS")
	poll_url := "https://api.openweathermap.org/data/2.5/weather?id=" + location_id + "&units=" + units + "&appid=" + api_key

	// Poll the API
	response, err := http.Get(poll_url)

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
	}
	
	// Show the data we got
    fmt.Println(string(responseData))

}

