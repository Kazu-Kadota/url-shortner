package encode

import (
	"fmt"

	convertasciitobase62 "github.com/Kazu-Kadota/url-shortner/utils/convertASCIIToBase62"
)

func Encode(input string) (encode []byte, err error) {
	data := []byte(input)
	var base62Data []byte
	for _, char := range data {
		base62number, _ := convertasciitobase62.Base62CharToValue(char)

		base62Data = append(base62Data, byte(base62number))
	}

	return base62Data, nil
}

func Decode(input []byte) (decode string, err error) {
	var base62Data []byte
	for _, integer := range input {
		base62char, _ := convertasciitobase62.Base62ValueToChar(integer)

		base62Data = append(base62Data, byte(base62char))
	}
	fmt.Println("DecodedByte: ", base62Data)

	if err != nil {
		return "", err
	}

	return string(base62Data), nil
}
