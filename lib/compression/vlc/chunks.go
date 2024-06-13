package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type encodingTable map[rune]string
type BinaryChunks []BinaryChunk
type BinaryChunk string

// Join joins chunks into string
func (bs BinaryChunks) Join() string {
	var buf strings.Builder

	for _, b := range bs {
		buf.WriteString(string(b))
	}
	return buf.String()
}

func (bs BinaryChunks) Bytes() []byte {
	res := make([]byte, 0, len(bs))
	for _, b := range bs {
		res = append(res, b.Byte())
	}
	return res
}

func (b BinaryChunk) Byte() byte {
	num, err := strconv.ParseUint(string(b), 2, chunkSize)
	if err != nil {
		panic("failedt to parse string to byte: " + err.Error())
	}
	return byte(num)
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

func NewBinChunks(data []byte) BinaryChunks {
	res := make(BinaryChunks, 0, len(data))

	for _, part := range data {
		res = append(res, NewBinChunk(part))
	}
	return res
}
func NewBinChunk(b byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", b))
}
