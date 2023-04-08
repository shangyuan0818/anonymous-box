package util

import "math/rand"

// RandNum returns a random number between min and max
func RandNum(min, max int) int {
	return min + rand.Intn(max-min)
}

// RandString returns a random string with the specified length
func RandString(length int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]byte, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}
