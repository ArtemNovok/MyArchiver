package shennofanno

import (
	"fmt"
	"math"
	"myarchiver/lib/compression/vlc/table"
	"sort"
	"strings"
)

type Generator struct {
}

type CharStat map[rune]int

func NewGenerator() Generator {
	return Generator{}
}

type encodingTable map[rune]code

type code struct {
	Char     rune
	Quantity int
	Bits     uint32
	Size     int
}

func (g Generator) NewTabel(text string) table.EncodingTable {
	// char  appearance
	stat := newCharStat(text)
	// encoding table
	table := build(stat)
	// return table.EncodingTable
	return table.Export()
}

func (t encodingTable) Export() map[rune]string {
	res := make(map[rune]string)
	for char, code := range t {
		byteStr := fmt.Sprintf("%b", code.Bits)
		if lenDiff := code.Size - len(byteStr); lenDiff > 0 {
			byteStr = strings.Repeat("0", lenDiff) + byteStr

		}
		res[char] = byteStr
	}
	return res
}

func build(stat CharStat) encodingTable {
	codes := make([]code, 0, len(stat))
	for char, quant := range stat {
		codes = append(codes, code{
			Char:     char,
			Quantity: quant,
		})
	}
	sort.Slice(codes, func(i, j int) bool {
		if codes[i].Quantity != codes[j].Quantity {
			return codes[i].Quantity > codes[j].Quantity
		}
		return codes[i].Char < codes[j].Char
	})

	assignCodes(codes)

	res := make(encodingTable)
	for _, code := range codes {
		res[code.Char] = code
	}

	return res
}

func assignCodes(codes []code) {
	if len(codes) < 2 {
		return
	}
	divider := bestDividerPosition(codes)
	for i := 0; i < len(codes); i++ {
		codes[i].Bits <<= 1
		codes[i].Size++
		if i >= divider {
			codes[i].Bits |= 1

		}
	}
	assignCodes(codes[:divider])
	assignCodes(codes[divider:])
}

func bestDividerPosition(codes []code) int {
	bestPosition := 0
	total := 0
	left := 0
	for _, code := range codes {
		total += code.Quantity
	}
	prefDiff := math.MaxInt
	for i := 0; i < len(codes)-1; i++ {
		left += codes[0].Quantity
		right := total - left
		diff := abs(right - left)
		if diff >= prefDiff {
			break
		}
		prefDiff = diff
		bestPosition = i + 1
	}
	return bestPosition
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func newCharStat(text string) CharStat {
	res := make(CharStat)
	for _, char := range text {
		res[char]++
	}
	return res
}
