package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name  string
	Craft string
}

type ResponseBody struct {
	Message string
	People  []Person
	// ^ Experiment to change this to []*Person
	Number int
}

// TODO: Refactor into reusable getAstros()

func main() {
	//
	client := http.Client{
		Timeout: 2 * time.Second, // Best practice
	}

	url := "http://api.open-notify.org/astros.json"

	// get is not valid, must get GET
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		// Must have better error handling
		log.Fatal("failed to create new req: ", err)
	}

	// req.Header.Set("User-Agent", "kid-server")
	res, err := client.Do(req)

	if err != nil {
		log.Fatal("failed to do req: ", err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Must understand the possible status code to build robust client
	// res.Status

	b, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal("failed to read res.Body")
	}

	// fmt.Printf("res.Body: %s", string(b))

	var resBody ResponseBody
	err = json.Unmarshal(b, &resBody)
	// ^ Question: Does this make sense to do this if status is not 200

	if err != nil {
		log.Fatalf("failed to unmarshal resBody %q error %s\n", string(b), err.Error())
	}

	fmt.Printf("res.Body: %+v\n", resBody)
	fmt.Printf("res.Body.People[0]: %+v\n", resBody.People[0])

}
