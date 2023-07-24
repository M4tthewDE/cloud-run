package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	command := Command{Command: os.Args[1]}
	data, err := json.Marshal(command)
	if err != nil {
		log.Fatalln(err)
	}

	request, err := http.NewRequest("POST", "http://localhost:8080/cmd", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalln(err)
	}

	request.Header.Add("TOKEN", os.Getenv("TOKEN"))

	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != 200 {
		log.Fatalln(resp.StatusCode)
	}

	var result Result
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Stdout: %s\n", result.Stdout)
	log.Printf("Stderr: %s\n", result.Stderr)
}

type Command struct {
	Command string
}

type Result struct {
	Stdout string
	Stderr string
}
