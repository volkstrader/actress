package actress

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProfile_Dating(t *testing.T) {
	a := Profile{
		Name: "test",
		Age:  25,
	}

	assert.True(t, a.Dating())
	a.Age = 26
	assert.False(t, a.Dating())
}
