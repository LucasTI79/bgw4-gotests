package doubles_test

import (
	"potato/doubles"
	"potato/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchByPhone(t *testing.T) {
	// given
	stubSearchEngine := doubles.StubSearchEngine{}
	engine := service.NewEngine(&stubSearchEngine)
	expectedName := "fulano"

	// when
	result, err := engine.SearchByPhone("12345678912")

	// then
	assert.Nil(t, err)
	assert.Equal(t, expectedName, result)
}
