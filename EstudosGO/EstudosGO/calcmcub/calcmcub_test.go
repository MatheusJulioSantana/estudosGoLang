package calcmcub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstanceNew(t *testing.T) {
	cc, err := NewCalcMC(WithAltura(10), WithBase(10), WithLargura(10))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
}
func TestOptionsParamsAltura(t *testing.T) {
	cc, err := NewCalcMC(WithAltura(10))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	assert.Equal(t, 10.0, cc.GetAltura())
}
func TestOptionsParamsBase(t *testing.T) {
	cc, err := NewCalcMC(WithBase(10))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	assert.Equal(t, 10.0, cc.GetBase())
}
func TestOptionsParamsLargura(t *testing.T) {
	cc, err := NewCalcMC(WithLargura(10))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	assert.Equal(t, 10.0, cc.GetLargura())
}
func TestCalculaMC(t *testing.T) {
	cc, err := NewCalcMC(WithAltura(10), WithBase(10), WithLargura(10))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	assert.Equal(t, 1000.0, cc.CalculaMC())
}
func TestCalculaMCWithError(t *testing.T) {
	cc, err := NewCalcMC(WithAltura(0), WithBase(10), WithLargura(10))
	assert.Error(t, err)
	assert.Nil(t, cc)

	cc, err = NewCalcMC(WithAltura(10), WithBase(0), WithLargura(0))
	assert.Error(t, err)
	assert.Nil(t, cc)
}
func TestSetAltura(t *testing.T) {
	cc, err := NewCalcMC(WithAltura(10), WithBase(10), WithLargura(10))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	cc.SetAltura(20)
	assert.Equal(t, 20.0, cc.GetAltura())
}
func TestSetLargura(t *testing.T) {
	cc, err := NewCalcMC(WithAltura(10), WithBase(10), WithLargura(10))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	cc.SetLargura(20)
	assert.Equal(t, 20.0, cc.GetLargura())
}
func TestSetBase(t *testing.T) {
	cc, err := NewCalcMC(WithAltura(10), WithBase(10), WithLargura(10))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	cc.SetBase(20)
	assert.Equal(t, 20.0, cc.GetBase())
}
