package n1n2perc

type CalcPerc struct {
	params *Params
}

func NewCalcPerc(opts ...Options) (*CalcPerc, error) {
	p, err := newParams(opts...)
	if err != nil {
		return nil, err
	}
	return &CalcPerc{
		params: p,
	}, nil
}
func (c *CalcPerc) CalculaPerc() float64 {
	return (c.params.GetN1() * c.params.GetN2()) * 10 * c.params.GetMulti()
}
func (c *CalcPerc) GetN1() float64 {
	return c.params.GetN1()
}
func (c *CalcPerc) GetN2() float64 {
	return c.params.GetN2()
}
func (c *CalcPerc) GetMulti() float64 {
	return c.params.GetMulti()
}
func (c *CalcPerc) SetN1(n1 float64) {
	c.params.SetN1(n1)
}
func (c *CalcPerc) SetN2(n2 float64) {
	c.params.SetN2(n2)
}
func (c *CalcPerc) SetMulti(multi float64) {
	c.params.SetMulti(multi)
}
