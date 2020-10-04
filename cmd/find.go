package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"goofi/cmd/cli"
	csv2 "goofi/cmd/csv"
	"goofi/cmd/output"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:     "find",
	Short:   "Look up abbreviation",
	Aliases: []string{"f", "finder", "wtf"},
	RunE: func(cmd *cobra.Command, args []string) error {
		arg := cli.FlatArgs(args)

		//return findOnline(arg)

		return findLocalResult(arg)
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
}

func findLocalResult(arg string) error {
	data, err := csv2.OpenAndReadCSV(goofiFilePath)
	if err != nil {
		return err
	}

	err = findAndPrintResult(data, arg)
	if err != nil {
		return err
	}

	return nil
}

func findOnline(arg string) error {
	userID := os.Getenv("abbreviations_userid")
	tokenID := os.Getenv("abbreviations_tokenid")
	if tokenID == "" || userID == "" {
		return errors.New("did not find user and token for abbreviations.com api")
	}

	url := fmt.Sprintf("https://www.stands4.com/services/v2/abbr.php?uid=%v&tokenid=%v&term=%v&format=json&sortby=p", userID, tokenID, arg)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	res := Results{}
	err = json.Unmarshal(respBody, &res)
	if err != nil {
		return err
	}

	maxRes := 5
	var data []string
	for i, r := range res.Result {
		if i == maxRes {
			break
		}
		data = append(data, r.Definition)
	}

	output.PrintResults(arg, data)

	return nil
}

type Results struct {
	Result []Result `json:"result"`
}

type Result struct {
	Term       string `json:"term"`
	Definition string `json:"definition"`
	Category   string `json:"category"`
}

func findAndPrintResult(data [][]string, arg string) error {
	var res []string
	for _, line := range data {
		if strings.ToLower(line[0]) == arg {
			res = append(res, line[1])
		}
	}

	output.PrintResults(arg, res)

	return nil
}
