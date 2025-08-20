package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghjklmnopqrstuwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generate a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for range n {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generate a random owner
func RandomOwner() string {
	return RandomString(6)
}

// Generate a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Generate a random currency code
func RandomCurrency() string {
	currencies := []string{USD, EUR, CAD}

	n := len(currencies)

	return currencies[rand.Intn(n)]
}
