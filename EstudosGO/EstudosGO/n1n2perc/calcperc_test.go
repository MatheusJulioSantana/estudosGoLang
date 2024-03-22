package n1n2perc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstanceNEw(t *testing.T) {
	cc, err := NewCalcPerc(WithN1(10.0), WithN2(4.0), WithMulti(0.12))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
}
func TestOptionsParamsN1(t *testing.T) {
	cc, err := NewCalcPerc(WithN1(10.0))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	assert.Equal(t, 10.0, cc.GetN1())

}
func TestOptionsParamsN2(t *testing.T) {
	cc, err := NewCalcPerc(WithN2(10.0))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	assert.Equal(t, 10.0, cc.GetN2())

}
func TestOptionsParamsMulti(t *testing.T) {
	cc, err := NewCalcPerc(WithMulti(0.12))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	assert.Equal(t, 0.12, cc.GetMulti())
}
func TestCalculaPerc(t *testing.T) {
	cc, err := NewCalcPerc(WithN1(10), WithN2(4), WithMulti(0.12))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	assert.Equal(t, 48.00, cc.CalculaPerc())
}
func TestCalculaPercWithError(t *testing.T) {
	cc, err := NewCalcPerc(WithN1(0), WithN2(4), WithMulti(0.12))
	assert.Error(t, err)
	assert.Nil(t, cc)

	cc, err = NewCalcPerc(WithN1(10), WithN2(0), WithMulti(0.12))
	assert.Error(t, err)
	assert.Nil(t, cc)

	cc, err = NewCalcPerc(WithN1(10), WithN2(0), WithMulti(0))
	assert.Error(t, err)
	assert.Nil(t, cc)
	assert.Nil(t, cc)
}
func TestSetN1(t *testing.T) {
	cc, err := NewCalcPerc(WithN1(10), WithN2(4), WithMulti(0.12))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	cc.SetN1(20)
	assert.Equal(t, 20.0, cc.GetN1())
}
func TestSetN2(t *testing.T) {
	cc, err := NewCalcPerc(WithN1(10), WithN2(4), WithMulti(0.12))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	cc.SetN2(20)
	assert.Equal(t, 20.0, cc.GetN2())
}
func TestSetMulti(t *testing.T) {
	cc, err := NewCalcPerc(WithN1(10), WithN2(10), WithMulti(0.10))
	assert.NoError(t, err)
	assert.NotNil(t, cc)
	cc.SetMulti(0.12)
	assert.Equal(t, 0.12, cc.GetMulti())

}
