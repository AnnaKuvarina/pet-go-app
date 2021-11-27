package utils

import "math/rand"

var letters = []rune("a@bcd12EefghSijklmHn$opqrs$gstu%&vwGDxyzAjsdd6m")

func GenerateRandomString(n uint32) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
