package cmd

import (
	"errors"
	"io"
	"myarchiver/lib/vlc"
	"os"

	"github.com/spf13/cobra"
)

var vlcUnpackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "unpack file using vlc algorithm",
	Run:   unpack,
}

const (
	unpackedFileExtension = "txt"
)

var (
	ErrEmptyPath = errors.New("path to file is not specified")
)

func unpack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		HandleErr(ErrEmptyPath)
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
	unpacked := vlc.Decode(data)
	err = os.WriteFile(FileName(filePath, unpackedFileExtension), []byte(unpacked), 0644)
	if err != nil {
		HandleErr(err)
	}

}
func init() {
	unpackCmd.AddCommand(vlcUnpackCmd)
}
