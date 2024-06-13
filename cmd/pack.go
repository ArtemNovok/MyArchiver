package cmd

import (
	"fmt"
	"io"
	"myarchiver/lib/compression"
	"myarchiver/lib/compression/vlc"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}

const packedExtension = "vlc"

func pack(cmd *cobra.Command, args []string) {
	var enc compression.Encode
	if len(args) == 0 || args[0] == "" {
		HandleErr(ErrEmptyPath)
	}
	method := cmd.Flag("method").Value.String()
	switch method {
	case "vlc":
		enc = vlc.New()
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
	packed := enc.Encode(string(data))
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
	rootCmd.AddCommand(packCmd)

	packCmd.Flags().StringP("method", "m", "", "compression method (vlc)")

	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
