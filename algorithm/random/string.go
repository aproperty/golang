package randomutil

import (
	"bytes"
	"math/rand"
	"time"
)

// Random generate string
func GetRandomNumString(n int) string {

	const alphanum = "012356789"
	lenAl := len(alphanum)

	var bytes = make([]byte, n)
	rand.Read(bytes)

	for i, bItem := range bytes {
		tm := bItem % byte(lenAl)
		bytes[i] = alphanum[tm]
	}

	return string(bytes)
}

func GenerateRandomString(len int64) string {

	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	strBytes := []byte(str)

	result := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var i int64
	for i = 0; i < len; i++ {
		idx := r.Intn(bytes.Count(strBytes, nil) - 1)
		temp := strBytes[idx]
		result = append(result, temp)
	}

	return string(result)
}
