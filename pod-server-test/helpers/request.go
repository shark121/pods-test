package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	// "github.com/pod-server-test/types"
)

func SendRides[T any](rides T, wg *sync.WaitGroup) {
	defer wg.Done()

	url := "http://localhost:8080/"

	client := &http.Client{}

	print(rides)

	data, err := json.Marshal(rides)

	// fmt.Println(string(data))

	if err != nil {
		print(err)
	}

	payload := bytes.NewBuffer(data)

	if err != nil {
		print(err)
	}

	// body := bytes.NewBuffer(payload)

	request, err := http.NewRequest("POST", url, payload)

	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		print(err)
	}

	resp, err := client.Do(request)

	if err != nil {
		fmt.Println("there was an error", err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		print(err)
	}

	fmt.Println(string(body), "######################")
}
