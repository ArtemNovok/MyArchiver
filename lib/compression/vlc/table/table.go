package table

import "strings"

type Generator interface {
	NewTable(text string) EncodingTable
}

type decodingTree struct {
	Value string
	Zero  *decodingTree
	One   *decodingTree
}

type EncodingTable map[rune]string

func (et EncodingTable) Decode(str string) string {
	dt := et.DecodingTree()

	return dt.Decode(str)
}

func (dt *decodingTree) Decode(str string) string {
	var buf strings.Builder
	currNode := dt

	for _, ch := range str {
		if currNode.Value != "" {
			buf.WriteString(currNode.Value)
			currNode = dt
		}
		switch ch {
		case '0':
			currNode = currNode.Zero
		case '1':
			currNode = currNode.One
		}
	}
	if currNode.Value != "" {
		buf.WriteString(currNode.Value)
		currNode = dt
	}
	return buf.String()

}

func (et EncodingTable) DecodingTree() decodingTree {
	res := decodingTree{}
	for ch, code := range et {
		res.add(code, ch)
	}
	return res
}

func (dt *decodingTree) add(code string, val rune) {
	curr := dt
	for _, ch := range code {
		switch ch {
		case '0':
			if curr.Zero == nil {
				curr.Zero = &decodingTree{}
			}
			curr = curr.Zero
		case '1':
			if curr.One == nil {
				curr.One = &decodingTree{}
			}
			curr = curr.One
		}
	}
	curr.Value = string(val)
}
