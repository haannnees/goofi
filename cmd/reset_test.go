package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveGoofiFile(t *testing.T) {
	assert.Error(t, removeGoofiFile(""))
}