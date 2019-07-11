package utils

import "log"

func ResponseError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
