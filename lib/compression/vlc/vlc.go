package vlc

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"myarchiver/lib/compression/vlc/table"
	"strings"
	"unicode"
)

const (
	chunkSize = 8
)

type EncoderDecoder struct {
	tblGenerator table.Generator
}

func New(generator table.Generator) EncoderDecoder {
	return EncoderDecoder{tblGenerator: generator}
}

// Encode encodes string using vlc algorithm
func (ed EncoderDecoder) Encode(str string) []byte {
	tbl := ed.tblGenerator.NewTable(str)
	// encoding to binary
	binStr := encodeBin(str, tbl)
	// bytes to hex and return
	return buildEncodedFile(tbl, binStr)
}

func (ed EncoderDecoder) Decode(dataEncoded []byte) string {
	tbl, data := parseFile(dataEncoded)
	return tbl.Decode(data)
}

func parseFile(data []byte) (table.EncodingTable, string) {
	const (
		tableSizeBinaryCount = 4
		dataSizeBinaryCount  = 4
	)
	tableSizeBin, data := data[:tableSizeBinaryCount], data[tableSizeBinaryCount:]
	dataSizeBin, data := data[:dataSizeBinaryCount], data[dataSizeBinaryCount:]
	tableSize := binary.BigEndian.Uint32(tableSizeBin)
	dataSize := binary.BigEndian.Uint32(dataSizeBin)

	tableBin, data := data[:tableSize], data[tableSize:]
	tbl := decodeTable(tableBin)
	body := NewBinChunks(data).Join()
	return tbl, body[:dataSize]
}

func decodeTable(tblBin []byte) table.EncodingTable {
	var tbl table.EncodingTable
	r := bytes.NewReader(tblBin)
	if err := gob.NewDecoder(r).Decode(&tbl); err != nil {
		panic("failed to decode table " + err.Error())
	}
	return tbl
}
func buildEncodedFile(tbl table.EncodingTable, data string) []byte {
	encodedTable := encodeTable(tbl)
	var buf bytes.Buffer
	buf.Write(encodeInt(len(encodedTable)))
	buf.Write(encodeInt(len(data)))
	buf.Write(encodedTable)
	buf.Write(splitByChunks(data, chunkSize).Bytes())
	return buf.Bytes()
}

func encodeInt(num int) []byte {

	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(num))
	return res
}
func encodeTable(tbl table.EncodingTable) []byte {
	var tableBuf bytes.Buffer
	err := gob.NewEncoder(&tableBuf).Encode(tbl)
	if err != nil {
		panic("can't serialize table " + err.Error())
	}
	return tableBuf.Bytes()
}

// encodeBin encodes string into string without spaces
// and with only 0 and 1
func encodeBin(str string, tbl table.EncodingTable) string {
	var buf strings.Builder
	for _, char := range str {
		buf.WriteString(bin(char, tbl))
	}
	return buf.String()
}

// bin transforms rune into bit string using table from getEncodingTable
func bin(r rune, tbl table.EncodingTable) string {
	res, ok := tbl[r]
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
