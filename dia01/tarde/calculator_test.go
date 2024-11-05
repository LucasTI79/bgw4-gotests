package calculator_test

import (
	calculator "potato"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDivision(t *testing.T) {
	// Organizar
	num1 := 3
	num2 := 0
	expectedResult := 0

	// Atuação
	result, err := calculator.Division(num1, num2)

	// Afirmação
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result, "must be equal")
}

func TestAdd(t *testing.T) {
	// Dado
	num1 := 3
	num2 := 5
	expectedResult := 8

	// Quando
	result, err := calculator.Add(num1, num2)

	// Então
	require.Nil(t, err)
	require.Equal(t, expectedResult, result, "must be equal")
}
