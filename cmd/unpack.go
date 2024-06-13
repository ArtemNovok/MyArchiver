package cmd

import (
	"errors"
	"io"
	"myarchiver/lib/compression"
	"myarchiver/lib/compression/vlc"
	shennofanno "myarchiver/lib/compression/vlc/table/shenno_fanno"
	"os"

	"github.com/spf13/cobra"
)

var unpackCmd = cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
	Run:   unpack,
}

const (
	unpackedFileExtension = "txt"
)

var (
	ErrEmptyPath = errors.New("path to file is not specified")
)

func unpack(cmd *cobra.Command, args []string) {
	var dec compression.Decode
	if len(args) == 0 || args[0] == "" {
		HandleErr(ErrEmptyPath)
	}
	method := cmd.Flag("method").Value.String()
	switch method {
	case "vlc":
		dec = vlc.New(shennofanno.NewGenerator())
	default:
		cmd.PrintErr("unknown method")
	}
	filePath := args[0]
	r, err := os.Open(filePath)
	if err != nil {
		HandleErr(err)
	}
	defer r.Close()
	data, err := io.ReadAll(r)
	if err != nil {
		HandleErr(err)
	}
	unpacked := dec.Decode(data)
	err = os.WriteFile(FileName(filePath, unpackedFileExtension), []byte(unpacked), 0644)
	if err != nil {
		HandleErr(err)
	}

}
func init() {
	rootCmd.AddCommand(&unpackCmd)
	unpackCmd.Flags().StringP("method", "m", "", "decompression method (vlc)")

	if err := unpackCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
