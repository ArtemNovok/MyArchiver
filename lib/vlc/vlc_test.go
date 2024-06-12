package vlc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrepareText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "happy path",
			str:  "hi",
			want: "hi",
		},
		{
			name: "one upper case",
			str:  "Hi",
			want: "!hi",
		},
		{
			name: "multiple upper cases",
			str:  "Hi my name is SomeName",
			want: "!hi my name is !some!name",
		},
	}
	for _, test := range tests {
		res := prepareText(test.str)
		require.Equal(t, res, test.want)
	}
}

func TestEncodeBin(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "happy path",
			str:  "!ted",
			want: "001000100110100101",
		},
	}
	for _, test := range tests {
		res := encodeBin(test.str)
		require.Equal(t, res, test.want)
	}
}

func TestSpitByChunks(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}

	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "happy path",
			args: args{
				bStr:      "001000100110100101",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100010", "01101001", "01000000"},
		},
	}
	for _, test := range tests {
		res := splitByChunks(test.args.bStr, test.args.chunkSize)
		require.Equal(t, res, test.want)
	}

}

func TestToHex(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want HexChunks
	}{
		{
			name: "happy path",
			bcs:  BinaryChunks{"0101111", "10000000"},
			want: HexChunks{"2F", "80"},
		},
	}
	for _, test := range tests {
		res := test.bcs.ToHex()
		require.Equal(t, res, test.want)
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "happy path",
			str:  "My name is Ted",
			want: "20 30 3C 18 77 4A E4 4D 28",
		},
	}
	for _, test := range tests {
		res := Encode(test.str)
		require.Equal(t, res, test.want)
	}
}
