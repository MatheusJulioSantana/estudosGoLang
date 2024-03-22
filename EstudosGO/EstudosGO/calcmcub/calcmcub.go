package calcmcub

type CalMetroCubico struct {
	params *Params
}

func NewCalcMC(opts ...Options) (*CalMetroCubico, error) {
	p, err := newParams(opts...)
	if err != nil {
		return nil, err
	}
	return &CalMetroCubico{
		params: p,
	}, nil
}
func (c *CalMetroCubico) CalculaMC() float64 {
	return c.params.GetALtura() * c.params.GetBase() * c.params.GetLargura()
}
func (c *CalMetroCubico) GetAltura() float64 {
	return c.params.GetALtura()
}
func (c *CalMetroCubico) GetBase() float64 {
	return c.params.GetBase()
}
func (c *CalMetroCubico) GetLargura() float64 {
	return c.params.GetLargura()
}

func (c *CalMetroCubico) SetAltura(altura float64) {
	c.params.SetAltura(altura)
}
func (c *CalMetroCubico) SetBase(base float64) {
	c.params.SetBase(base)
}
func (c *CalMetroCubico) SetLargura(largura float64) {
	c.params.SetLargura(largura)
}
