package cmd

import (
	"bufio"
	"fmt"
	csv2 "goofi/cmd/csv"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:     "clean",
	Short:   "Clean up the local glossary (remove duplicates and sort)",
	Aliases: []string{"c"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return sortAndRemoveDuplicates()
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

//TODO refactor
func sortAndRemoveDuplicates() error {
	file, err := os.Open(goofiFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = file.Close()
	if err != nil {
		return err
	}

	err = os.Remove(goofiFilePath)
	if err != nil {
		return err
	}

	err = csv2.CreateCsvFile(goofiFilePath)
	if err != nil {
		return err
	}

	newFile, err := os.OpenFile(goofiFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer newFile.Close()

	sortSlice(lines)

	occurred := make(map[string]bool, len(lines))
	for e := range lines {
		if occurred[lines[e]] != true {
			occurred[lines[e]] = true
			_, err = newFile.WriteString(lines[e])
			if err != nil {
				return err
			}
			_, err = newFile.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}

	fmt.Print("Successfully cleaned goofi!\n")

	return nil
}

func sortSlice(lines []string) {
	sort.Slice(lines, func(i, j int) bool { return strings.Compare(lines[i], lines[j]) < 0 })
}
