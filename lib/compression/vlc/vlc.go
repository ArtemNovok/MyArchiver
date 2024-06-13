package vlc

import (
	"myarchiver/lib/compression/vlc/table"
	"strings"
	"unicode"
)

const (
	chunkSize = 8
)

type EncoderDecoder struct {
}

func New() EncoderDecoder {
	return EncoderDecoder{}
}

// Encode encodes string using vlc algorithm
func (_ EncoderDecoder) Encode(str string) []byte {

	// lower case with ! (M -> !m)
	str = prepareText(str)
	// encoding to binary
	binStr := encodeBin(str)

	// split bits to bytes (8)
	chunks := splitByChunks(binStr, chunkSize)

	// bytes to hex and return
	return chunks.Bytes()
}

// Decode decodes string that is product of vlc algorithm to text
func (_ EncoderDecoder) Decode(bytes []byte) string {
	bString := NewBinChunks(bytes).Join()
	// build decoding tree
	bTree := getEncodingTable().DecodingTree()
	// convert binary string to usual string
	decodedStr := bTree.Decode(bString)
	// return decoded string
	return exportText(decodedStr)
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
func getEncodingTable() table.EncodingTable {
	return table.EncodingTable{
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

// exportText export text ("!my name is !some!name" -> "My name is SomeName")
func exportText(str string) string {
	var buf strings.Builder
	var Capital bool
	for _, ch := range str {
		if ch == '!' {
			Capital = true
			continue
		}
		if Capital {
			buf.WriteRune(unicode.ToUpper(ch))
			Capital = false
		} else {
			buf.WriteRune(ch)
		}
	}
	return buf.String()
}
