package doubles_test

import (
	"potato/doubles"
	"potato/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchByName(t *testing.T) {
	// given
	searchEngine := doubles.MockSearchEngine{}
	engine := service.NewEngine(&searchEngine)
	expectedResultSearchByPhone := "fulano"

	// when
	engine.SearchByName("fulano")
	resultSearchByPhone, err := engine.SearchByPhone("12345678912")

	// then
	assert.Nil(t, err)
	assert.Equal(t, expectedResultSearchByPhone, resultSearchByPhone)
	assert.True(t, searchEngine.SearchByNameWasCalled)
}
