package common

import (
	"log"
)

func LogErrorService(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func LogSuccess(msg string) {
	log.Printf("SUCCESS: %s", msg)
}