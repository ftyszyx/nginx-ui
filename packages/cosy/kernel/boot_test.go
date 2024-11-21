package kernel

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var a = 1

func TestBoot(t *testing.T) {
	RegisterAsyncFunc(func() {
		a = 2
	})
	RegisterSyncsFunc(func() {
		a = 3
	})
	Boot()
	time.Sleep(1 * time.Second)

	assert.Equal(t, 2, a, "a should be 3")
}
