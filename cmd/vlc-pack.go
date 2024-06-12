package cmd

import (
	"fmt"
	"io"
	"myarchiver/lib/vlc"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const (
	packedExtension = "vlc"
)

var vlcPackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

func pack(_ *cobra.Command, args []string) {
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
	packed := vlc.Encode(string(data))
	err = os.WriteFile(FileName(filePath, packedExtension), packed, 0644)
	if err != nil {
		HandleErr(err)
	}

}

// FileName returns file name from given path and with given extension
func FileName(path string, ext string) string {
	fileName := filepath.Base(path)
	return fmt.Sprintf("%s.%s", strings.TrimSuffix(fileName, filepath.Ext(fileName)), ext)
}

func init() {
	packCmd.AddCommand(vlcPackCmd)
}
