package doubles_test

import (
	"potato/doubles"
	"potato/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerchByPhone(t *testing.T) {
	// given
	spySearchEngine := doubles.SpySearchEngine{}
	engine := service.NewEngine(&spySearchEngine)

	// when

	_, err := engine.SearchByPhone("12345678912")
	assert.Nil(t, err)
	assert.True(t, spySearchEngine.SearchByPhoneWasCalled)

}
