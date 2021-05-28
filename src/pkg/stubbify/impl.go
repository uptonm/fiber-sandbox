package stubbify

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	base uint64 = 62
	charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

// Base62 encoding from this article https://medium.com/analytics-vidhya/base-62-text-encoding-decoding-b43921c7a954

func BaseEncode(num uint64) string {
	if num == 0 {
		return string(charset[0])
	}
	encoded := ""
	for num > 0 {
		r := num % base
		num /= base
		encoded = string(charset[r]) + encoded

	}
	return encoded
}

func BaseDecode(encoded string) (uint64, error) {
	var val uint64
	for index, char := range encoded {
		pow := len(encoded) - (index + 1)
		pos := strings.IndexRune(charset, char)
		if pos == -1 {
			return 0, errors.New("invalid character: " + string(char))
		}

		val += uint64(pos) * uint64(math.Pow(float64(base), float64(pow)))
	}

	return val, nil
}

const encodingChunkSize = 2

// no of bytes required in base62 to represent hex encoded string value of length encodingChunkSize
// given by formula :: int(math.Ceil(math.Log(math.Pow(16, 2*encodingChunkSize)-1) / math.Log(62)))
const decodingChunkSize = 3

func Encode(str string) string {
	var encoded strings.Builder

	inBytes := []byte(str)
	byteLength := len(inBytes)

	for i := 0; i < byteLength; i += encodingChunkSize {
		chunk := inBytes[i:minOf(i+encodingChunkSize, byteLength)]
		s := hex.EncodeToString(chunk)
		val, _ := strconv.ParseUint(s, 16, 64)
		w := padLeft(BaseEncode(val), "0", decodingChunkSize)
		encoded.WriteString(w)
	}
	return encoded.String()
}

func Decode(encoded string) (string, error) {
	decodedBytes := []byte{}
	for i := 0; i < len(encoded); i += decodingChunkSize {
		chunk := encoded[i:minOf(i+decodingChunkSize, len(encoded))]
		val, err := BaseDecode(chunk)
		if err != nil {
			return "", err
		}
		chunkHex := strconv.FormatUint(val, 16)
		dst := make([]byte, hex.DecodedLen(len([]byte(chunkHex))))
		_, err = hex.Decode(dst, []byte(chunkHex))
		if err != nil {
			return "", errors.New(fmt.Sprintf("malformed input: %s", err.Error()))
		}
		decodedBytes = append(decodedBytes, dst...)
	}
	s := string(decodedBytes)
	return s, nil
}

func minOf(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func padLeft(str, pad string, length int) string {
	for len(str) < length {
		str = pad + str
	}
	return str
}