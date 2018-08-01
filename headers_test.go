package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Headers_Set_When_Receiver_Is_Nil(t *testing.T) {
	var h Headers
	assert.Nil(t, h)
	assert.Nil(t, h.Set("a:b"))
}

func Test_Headers_Set_WhenValueIsInCorrectFormat(t *testing.T) {
	var h Headers
	h.Set("a:b,c:d")
	assert.Contains(t, h, "a")
	assert.Contains(t, h, "c")
}

func Test_Headers_Set_WhenValueIsNotInCorrectFormat(t *testing.T) {
	var h Headers
	err := h.Set("a")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Invalid header")
}
