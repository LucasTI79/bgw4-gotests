package doubles_test

import (
	"potato/doubles"
	"potato/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchEngineVersion(t *testing.T) {
	// given
	searchEngine := doubles.DummySearchEngine{}
	engine := service.NewEngine(&searchEngine)
	expectedResult := "version: 1.0.0"

	// when
	result := engine.GetVersion()

	// then
	assert.Equal(t, expectedResult, result)
}
