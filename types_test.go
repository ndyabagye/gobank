package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccount(t *testing.T) {
	acc, err := NewAccount("a", "b", "pass1234")
	assert.Nil(t, err)

	fmt.Printf("%+v\n", acc)
}
