package utils

import "log"

func Close(closeFunc func() error) {
	if err := closeFunc(); err != nil {
		log.Fatal(err)
	}
}
