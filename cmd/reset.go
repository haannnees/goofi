package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:     "reset",
	Short:   "Remove all added abbreviations and reset goofi",
	Aliases: []string{"r"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return removeGoofiFile()
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}

func removeGoofiFile() error {
	err := os.Remove(goofiFilePath)

	if err != nil {
		return errors.New("could not reset goofi")
	}

	fmt.Print("Successfully reset goofi!\n")

	return nil
}
