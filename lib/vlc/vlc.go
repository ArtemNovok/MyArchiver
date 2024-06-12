package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type encodingTable map[rune]string
type BinaryChunks []BinaryChunk
type BinaryChunk string
type HexChunk string
type HexChunks []HexChunk

const (
	chunkSize = 8
)

func (bs BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(bs))
	for _, chunk := range bs {
		hChunk := chunk.ToHex()
		res = append(res, hChunk)
	}

	return res
}

func (b BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseUint(string(b), 2, chunkSize)
	if err != nil {
		panic("can't parse chunk " + err.Error())
	}
	res := strings.ToUpper(fmt.Sprintf("%x", num))
	if len(res) == 1 {
		res = "0" + res
	}
	return HexChunk(res)
}

func Encode(str string) string {
	// lower case with ! (M -> !m)
	str = prepareText(str)
	// encoding to binary
	binStr := encodeBin(str)

	// split bits to bytes (8)
	chunks := splitByChunks(binStr, chunkSize)

	// bytes to hex and return
	return string(chunks.ToHex().ToString())
}

func (hchs HexChunks) ToString() string {
	const separator = " "
	switch len(hchs) {
	case 0:
		return ""
	case 1:
		return string(hchs[0])
	}

	var buf strings.Builder
	buf.WriteString(string(hchs[0]))
	for _, hch := range hchs[1:] {
		buf.WriteString(separator)
		buf.WriteString(string(hch))
	}
	return buf.String()

}

// prepareText prepares text so all upper case letters
// are transformed to lower case with ! (P -> !p)
func prepareText(str string) string {
	var buf strings.Builder
	for _, char := range str {
		if unicode.IsUpper(char) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(char))
		} else {
			buf.WriteRune(char)
		}
	}
	return buf.String()
}

// encodeBin encodes string into string without spaces
// and with only 0 and 1
func encodeBin(str string) string {
	var buf strings.Builder
	for _, char := range str {
		buf.WriteString(bin(char))
	}
	return buf.String()
}

// bin transforms rune into bit string using table from getEncodingTable
func bin(r rune) string {
	table := getEncodingTable()
	res, ok := table[r]
	if !ok {
		panic("unknown character: " + string(r))
	}
	return res
}

// getEncodingTable returns encoding table
func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		't': "1001",
		'n': "10000",
		's': "0101",
		'r': "01000",
		'd': "00101",
		'!': "001000",
		'c': "000101",
		'm': "000011",
		'g': "0000100",
		'b': "0000010",
		'v': "00000001",
		'k': "0000000001",
		'q': "000000000001",
		'e': "101",
		'o': "10001",
		'a': "011",
		'i': "01001",
		'h': "0011",
		'l': "001001",
		'u': "00011",
		'f': "000100",
		'p': "0000101",
		'w': "0000011",
		'y': "0000001",
		'j': "000000001",
		'x': "00000000001",
		'z': "000000000000",
	}
}

// splitByChunks splits binary string by chunks with given size
func splitByChunks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)
	chunksCount := strLen / chunkSize
	if strLen/chunksCount != 0 {
		chunksCount++
	}
	res := make(BinaryChunks, 0, chunksCount)

	var buf strings.Builder

	for ind, char := range bStr {
		buf.WriteString(string(char))
		if (ind+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))
		res = append(res, BinaryChunk(lastChunk))
	}
	return res
}
