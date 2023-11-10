package utils

import (
	"log"
)

// Logs an error to the console
func LogError(errMsg string, err error) {
	log.Println("ERROR:", errMsg+":", err)
}

// Logs a success to the console
func LogSuccess(succMsg string) {
	log.Println("SUCCESS:", succMsg)
}

func LogFatal(fatalMsg string, err error) {
	log.Fatalln("FATAL:", fatalMsg+":", err)
}
