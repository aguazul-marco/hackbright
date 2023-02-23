package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string for password and email of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomUser generates a random owner
func RandomUser() string {
	return RandomString(6)
}

// RandomRating generates a random rating from 1 to 5
func RandomRating() int32 {
	return int32(RandomInt(1, 5))
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf(RandomString(24) + "@email.com")
}

func RandomURL() string {
	return fmt.Sprintf("www.randomurl.com//" + RandomString(24))
}

// func randate() time.Time {
// 	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
// 	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
// 	delta := max - min

// 	sec := rand.Int63n(delta) + min

// 	return time.Unix(sec, 0)
// }
