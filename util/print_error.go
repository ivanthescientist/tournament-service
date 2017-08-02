package util

import "log"

func PrintError(message string, err error)  {
	if err != nil {
		log.Printf("%s: %s", message, err.Error())
	}
}