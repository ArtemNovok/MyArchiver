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

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want []byte
	}{
		{
			name: "happy path",
			str:  "My name is Ted",
			want: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
		},
	}
	for _, test := range tests {
		res := Encode(test.str)
		require.Equal(t, res, test.want)
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		name string
		bchs BinaryChunks
		want string
	}{
		{
			name: "happy path",
			bchs: BinaryChunks{"00101111", "10000000"},
			want: "0010111110000000",
		},
	}
	for _, test := range tests {
		res := test.bchs.Join()
		require.Equal(t, res, test.want)
	}
}

func TestExportText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "Happy path",
			str:  "!my name is !some!name",
			want: "My name is SomeName",
		},
	}
	for _, test := range tests {
		res := exportText(test.str)
		require.Equal(t, res, test.want)
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want string
	}{
		{
			name: "Happy path",
			data: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
			want: "My name is Ted",
		},
	}
	for _, test := range tests {
		res := Decode(test.data)
		require.Equal(t, res, test.want)
	}
}

func TestNewBinChunks(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want BinaryChunks
	}{
		{
			name: "happy path",
			data: []byte{20, 30, 60, 18},
			want: BinaryChunks{"00010100", "00011110", "00111100", "00010010"},
		},
	}
	for _, test := range tests {
		res := NewBinChunks(test.data)
		require.Equal(t, res, test.want)
	}
}
