package doubles_test

import (
	"potato/doubles"
	"potato/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFakeSearchByPhone(t *testing.T) {
	// given
	searchEngine := doubles.MockSearchEngine{}
	engine := service.NewEngine(&searchEngine)
	expectedResult := "fulano"
	engine.AddEntry("fulano", "12345678912")

	// when
	result, err := engine.SearchByPhone("12345678912")

	// then
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}
