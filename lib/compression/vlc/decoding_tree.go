package vlc

import "strings"

type DecodingTree struct {
	Value string
	Zero  *DecodingTree
	One   *DecodingTree
}

func (dt *DecodingTree) Decode(str string) string {
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

func (et encodingTable) DecodingTree() DecodingTree {
	res := DecodingTree{}
	for ch, code := range et {
		res.Add(code, ch)
	}
	return res
}

func (dt *DecodingTree) Add(code string, val rune) {
	curr := dt
	for _, ch := range code {
		switch ch {
		case '0':
			if curr.Zero == nil {
				curr.Zero = &DecodingTree{}
			}
			curr = curr.Zero
		case '1':
			if curr.One == nil {
				curr.One = &DecodingTree{}
			}
			curr = curr.One
		}
	}
	curr.Value = string(val)
}
