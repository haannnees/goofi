package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func OpenAndReadCSV(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
		return nil, err
	}

	reader := csv.NewReader(file)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func OpenAndWriteCSV(data [][]string, path string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Print("Could not open file")
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(data)
	if err != nil {
		log.Print("Could not write to file")
		return err
	}

	return nil
}

func OpenAndWriteSingleLineToCSV(line []string, path string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Print("Could not open file")
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(line)
	if err != nil {
		log.Print("Could not write to file")
		return err
	}

	return nil
}

func OpenAndReadRemoteCSV(remoteUrl string) ([][]string, error) {
	_, err := url.ParseRequestURI(remoteUrl)
	if err != nil {
		return nil, errors.New("please enter valid url")
	}

	resp, err := http.Get(remoteUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func CreateCsvFile(path string) error {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	return err
}
