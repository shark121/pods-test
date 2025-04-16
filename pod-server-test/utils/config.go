package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type KEYS struct {
	Maps_key string `json:"GOOGLE_MAPS_API_KEY"`
}

func ReadConfig(path string) KEYS {
	file, err := os.Open(path)

	if err != nil {
		print(err)
	}

	bytes, err := io.ReadAll(file)

	var res KEYS

	err = json.Unmarshal(bytes, &res)

	if err != nil {
		fmt.Println("json unmarshalling error ", err)
	}

	return res

}
