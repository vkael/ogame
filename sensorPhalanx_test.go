package ogame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSensorPhalanxPrice(t *testing.T) {
	sp := newSensorPhalanx()
	assert.Equal(t, Resources{Metal: 20000, Crystal: 40000, Deuterium: 20000}, sp.GetPrice(1))
	assert.Equal(t, Resources{Metal: 40000, Crystal: 80000, Deuterium: 40000}, sp.GetPrice(2))
	assert.Equal(t, Resources{Metal: 80000, Crystal: 160000, Deuterium: 80000}, sp.GetPrice(3))
}
