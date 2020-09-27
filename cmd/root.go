package cmd

import (
	"errors"
	"fmt"
	"goofi/cmd/csv"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/mitchellh/go-homedir"
)

var (
	rootCmd = &cobra.Command{
		Use:   "goofi",
		Short: "CLI-Tool to look up abbreviations.",
		Long:  `With goofi you can find the meaning of abbreviations.`,
	}

	goofiFilePath = getGoofiFilePath()
)

//Execute initializes all commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	err := initGoofiFile()
	if err != nil {
		log.Fatal(err)
	}
}

//create local hidden csv file
func initGoofiFile() error {
	err := csv.CreateCsvFile(goofiFilePath)
	if err != nil {
		return errors.New("could not init goofi file")
	}

	return nil
}

//get path where goofi file will be stored
func getGoofiFilePath() string {
	home, err := homedir.Dir()
	if err != nil {
		user := os.Getenv("USER")
		return fmt.Sprintf("/Users/%s/.goofi.csv", user)
	}
	return fmt.Sprint(home + "/.goofi.csv")
}
