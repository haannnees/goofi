package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindOnline(t *testing.T) {
	assert.Error(t, findOnline(""))
}

func TestFindAndPrintResult(t *testing.T) {
	data := [][]string{
		{"CLI", "Command Line Interface"},
		{"A", "B"},
	}

	assert.Nil(t, findAndPrintResult(data, "CLI"))
	assert.Nil(t, findAndPrintResult(data, "A"))
}
