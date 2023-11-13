package utils

import (
	"log"
	"strconv"
)

// Logs an error to the console
func LogError(errMsg string, err error) {
	log.Println("ERROR:", errMsg+":", err)
}

func LogWarning(warnMsg string) {
	log.Println("WARNING:", warnMsg)
}

// Logs a success to the console
func LogSuccess(succMsg string) {
	log.Println("SUCCESS:", succMsg)
}

func LogFatal(fatalMsg string, err error) {
	log.Fatalln("FATAL:", fatalMsg+":", err)
}

func IntToString(integer int64) string {
	value := strconv.FormatInt(int64(integer), 10)

	return value
}
