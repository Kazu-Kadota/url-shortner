package main

import (
	"fmt"
	"math"

	base62 "github.com/Kazu-Kadota/url-shortner/utils/base62"
)

func main() {
	encode, _ := base62.Encode("Hello World!")

	var sumEncode float64 = 0
	for _, b := range encode {
		sumEncode += float64(b)
	}

	var decodedByte []byte
	var remainder int64
	var divider int64 = int64(sumEncode + math.Pow(62, 4))

	for divider >= 62 {

		remainder = divider % 62
		divider = int64(math.Round(float64(divider/62) - 0.5))

		decodedByte = append(decodedByte, byte(remainder))
	}

	decode, err := base62.Decode(decodedByte)

	if err != nil {
		fmt.Println("Error on decode: ", err)
	}

	fmt.Println("Decode: ", decode)
}
