package table

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodingTree(t *testing.T) {
	tests := []struct {
		name string
		et   EncodingTable
		want decodingTree
	}{
		{
			name: "base tree test",
			et: EncodingTable{
				'a': "11",
				'b': "1001",
				'z': "0101",
			},
			want: decodingTree{
				Zero: &decodingTree{
					One: &decodingTree{
						Zero: &decodingTree{
							One: &decodingTree{
								Value: "z",
							},
						},
					},
				},
				One: &decodingTree{
					Zero: &decodingTree{
						Zero: &decodingTree{
							One: &decodingTree{
								Value: "b",
							},
						},
					},
					One: &decodingTree{
						Value: "a",
					},
				},
			},
		},
	}
	for _, test := range tests {
		res := test.et.DecodingTree()
		require.Equal(t, res, test.want)
	}
}
