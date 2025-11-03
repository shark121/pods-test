package utils

import (
	"io"
	"log"
)

func HasErr(err error) {
	if err != nil && err != io.EOF {
		log.Fatal("there was an error: ", err)
	}
}
