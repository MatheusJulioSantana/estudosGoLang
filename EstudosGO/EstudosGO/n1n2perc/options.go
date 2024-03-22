package n1n2perc

import "errors"

type Options func(*Params) error

type Params struct {
	n1    float64
	n2    float64
	multi float64
}

func newParams(opts ...Options) (*Params, error) {
	p := &Params{}
	for _, opt := range opts {
		if err := opt(p); err != nil {
			return nil, err
		}
	}
	return p, nil
}

func WithN1(n1 float64) Options {
	return func(r *Params) error {
		if n1 <= 0 {
			return errors.New("n1 é 0")
		}
		r.n1 = n1
		return nil
	}
}
func WithN2(n2 float64) Options {
	return func(r *Params) error {
		if n2 <= 0 {
			return errors.New("n2 é 0")
		}
		r.n2 = n2
		return nil
	}
}
func WithMulti(multi float64) Options {
	return func(r *Params) error {
		if multi <= 0 {
			return errors.New("O multiplicador é 0")
		}
		r.multi = multi
		return nil
	}
}

func (p Params) GetN1() float64 {
	return p.n1
}
func (p Params) GetN2() float64 {
	return p.n2
}
func (p Params) GetMulti() float64 {
	return p.multi
}

func (p *Params) SetN1(n1 float64) {
	p.n1 = n1
}
func (p *Params) SetN2(n2 float64) {
	p.n2 = n2
}
func (p *Params) SetMulti(multi float64) {
	p.multi = multi
}
