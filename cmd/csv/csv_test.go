package csv

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAndDeleteCsvFile(t *testing.T) {
	dir, _ := homedir.Dir()
	path := fmt.Sprint(dir, "/test.csv")
	err := CreateCsvFile(path)
	assert.Nil(t, err)

	err = RemoveCSVFile(path)
	assert.Nil(t, err)
}
