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
type HexChunk string
type HexChunks []HexChunk

const separator = " "

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

func (hch HexChunk) ToBin() BinaryChunk {
	num, err := strconv.ParseUint(string(hch), 16, chunkSize)
	if err != nil {
		panic("failed to convert HexChunk to BinChunk: " + err.Error())
	}
	res := fmt.Sprintf("%08b", num)
	return BinaryChunk(res)
}

func (hchs HexChunks) ToBin() BinaryChunks {
	res := make(BinaryChunks, 0, len(hchs))
	for _, hch := range hchs {
		res = append(res, hch.ToBin())
	}
	return res
}

func (hchs HexChunks) ToString() string {
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

// Join joins chunks into string
func (bs BinaryChunks) Join() string {
	var buf strings.Builder

	for _, b := range bs {
		buf.WriteString(string(b))
	}
	return buf.String()
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

func NewHexChunks(str string) HexChunks {
	parts := strings.Split(str, separator)
	res := make(HexChunks, 0, len(parts))

	for _, part := range parts {
		res = append(res, HexChunk(part))
	}
	return res
}
