package utils

import "log"

func HandleError(err error) {
	if err != nil {
		log.Printf("Error: %s\n", err)
	}
}
