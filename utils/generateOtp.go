package utils

import (
	"math/rand"
)

func GenerateOTP() int {
	//rand.Seed(time.Now().UnixNano())
	return rand.Intn(9000) + 1000 // Ensures the number is 4 digits (between 1000 and 9999)
}
