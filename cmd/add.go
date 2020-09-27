package cmd

import (
	"errors"
	"fmt"
	"goofi/cmd/cli"
	csv2 "goofi/cmd/csv"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add new abbreviations to goofi",
	Aliases: []string{"a"},
	RunE: func(cmd *cobra.Command, args []string) error {
		singleAbbr, _ := cmd.Flags().GetBool("single")
		remoteAbbr, _ := cmd.Flags().GetBool("remote")
		localAbbr, _ := cmd.Flags().GetBool("local")

		arg := cli.FlatArgsSimple(args)

		if singleAbbr {
			return addSingleAbbreviation(arg)
		}

		if remoteAbbr {
			return addRemoteAbbreviations(arg)
		}

		if localAbbr {
			return addLocalAbbreviations(arg)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("single", "s", false, "Add single abbreviation to goofi (abbreviation=value)")
	addCmd.Flags().BoolP("remote", "r", false, "Add abbreviations from remote CSV file (e.g. Github)")
	addCmd.Flags().BoolP("local", "l", false, "Add abbreviations from local CSV file")
}

func addRemoteAbbreviations(remoteUrl string) error {
	data, err := csv2.OpenAndReadRemoteCSV(remoteUrl)
	if err != nil {
		fmt.Print("Cant read csv from url")
		return err
	}

	err = csv2.OpenAndWriteCSV(data, goofiFilePath)
	if err != nil {
		return err
	}

	fmt.Print("Successfully added remote CSV abbreviation list to goofi!\n")

	return err
}

func addLocalAbbreviations(filePath string) error {
	data, err := csv2.OpenAndReadCSV(filePath)
	if err != nil {
		return err
	}

	err = csv2.OpenAndWriteCSV(data, goofiFilePath)
	if err != nil {
		return err
	}

	fmt.Print("Successfully added local CSV abbreviation list to goofi!\n")

	return nil
}

func addSingleAbbreviation(arg string) error {
	splits := strings.Split(arg, "=")

	if len(splits) != 2 {
		return errors.New("abbreviation not correct formatted")
	}

	err := csv2.OpenAndWriteSingleLineToCSV([]string{splits[0], splits[1]}, goofiFilePath)
	if err != nil {
		return err
	}

	fmt.Print("Successfully added single abbreviation to goofi!\n")

	return nil
}
