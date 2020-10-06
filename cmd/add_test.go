package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddSingleAbbreviation(t *testing.T) {
	err := addSingleAbbreviation("abbreviation;meaning")
	assert.Error(t, err)
}