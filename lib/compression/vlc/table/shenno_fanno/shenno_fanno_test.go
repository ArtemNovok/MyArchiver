package shennofanno

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBestDividedPosition(t *testing.T) {
	tests := []struct {
		name string
		sl   []code
		want int
	}{
		{
			name: "happy path",
			sl:   []code{code{Char: 'a', Quantity: 2}},
			want: 0,
		},
		{
			name: "happy path2",
			sl:   []code{code{Char: 'a', Quantity: 2}, code{Char: 'b', Quantity: 2}},
			want: 1,
		},
		{
			name: "wrong res",
			sl:   []code{code{Char: 'a', Quantity: 2}, code{Char: 'b', Quantity: 2}},
			want: 3,
		},
		{
			name: "3 elements",
			sl:   []code{code{Char: 'a', Quantity: 5}, code{Char: 'b', Quantity: 2}, code{Char: 'c', Quantity: 3}},
			want: 1,
		},
		{
			name: "4 elements",
			sl: []code{code{Char: 'a', Quantity: 2}, code{Char: 'b', Quantity: 2},
				code{Char: 'c', Quantity: 1}, code{Char: 'g', Quantity: 1}},
			want: 1,
		},
	}
	for _, test := range tests {
		res := bestDividerPosition(test.sl)
		if test.name == "wrong res" {
			require.NotEqual(t, res, test.want)
		} else {
			require.Equal(t, res, test.want)
		}
	}
}

func TestAssignCodes(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  []code
	}{
		{
			name: "two elements",
			codes: []code{
				code{Quantity: 2},
				code{Quantity: 2},
			},
			want: []code{
				code{Quantity: 2, Bits: 0, Size: 1},
				code{Quantity: 2, Bits: 1, Size: 1},
			},
		},
		{
			name: "three elements, certain position",
			codes: []code{
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1}, // 0
				{Quantity: 1, Bits: 2, Size: 2}, // 10
				{Quantity: 1, Bits: 3, Size: 2}, // 11
			},
		},
		{
			name: "three elements, uncertain position",
			codes: []code{
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: []code{
				{Quantity: 1, Bits: 0, Size: 1}, // 0
				{Quantity: 1, Bits: 2, Size: 2}, // 10
				{Quantity: 1, Bits: 3, Size: 2}, // 11
			},
		},
	}
	for _, test := range tests {
		assignCodes(test.codes)
		require.Equal(t, test.codes, test.want)
	}

}

func TestBuild(t *testing.T) {
	tests := []struct {
		name string
		text string
		want encodingTable
	}{
		{
			name: "base test",
			text: "abbbcc",
			want: encodingTable{
				'a': code{
					Char:     'a',
					Quantity: 1,
					Bits:     3,
					Size:     2,
				},
				'b': code{
					Char:     'b',
					Quantity: 3,
					Bits:     0,
					Size:     1,
				},
				'c': code{
					Char:     'c',
					Quantity: 2,
					Bits:     2,
					Size:     2,
				},
			},
		},
		{
			name: "base test",
			text: "aaccbb",
			want: encodingTable{
				'a': code{
					Char:     'a',
					Quantity: 2,
					Bits:     0,
					Size:     1,
				},
				'b': code{
					Char:     'b',
					Quantity: 2,
					Bits:     2,
					Size:     2,
				},
				'c': code{
					Char:     'c',
					Quantity: 2,
					Bits:     3,
					Size:     2,
				},
			},
		},
	}
	for _, test := range tests {
		table := build(newCharStat(test.text))
		require.Equal(t, table, test.want)
	}
}
