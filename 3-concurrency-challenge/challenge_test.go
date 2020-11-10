package challenge

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

// func TestEqualOutput(t *testing.T) {
// 	sequential := CountLetterFrequency(one + two + three)
// 	concurrent := ConcurrentLetterFrequency([]string{one, two, three})

// 	if !reflect.DeepEqual(sequential, concurrent) {
// 		t.Fatal("These are not the same")
// 	}
// }

func generateString() []string {
	var completeString []string
	for i := 0; i < 100000; i++ {
		completeString = append(completeString, String(5))
	}
	return completeString
}
func BenchmarkSequential(b *testing.B) {
	inputString := generateString()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountLetterFrequency(strings.Join(inputString, ""))
	}
}

func BenchmarkConcurrent(b *testing.B) {
	inputString := generateString()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcurrentLetterFrequency(inputString)
	}
}
