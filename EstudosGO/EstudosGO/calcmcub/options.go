package calcmcub

import "errors"

type Options func(*Params) error

type Params struct {
	altura  float64
	base    float64
	largura float64
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

func WithAltura(altura float64) Options {
	return func(r *Params) error {
		if altura <= 0 {
			return errors.New("Altura Zerada")
		}
		r.altura = altura
		return nil
	}

}
func WithBase(base float64) Options {
	return func(r *Params) error {
		if base <= 0 {
			return errors.New("Base Zerada")
		}
		r.base = base
		return nil
	}
}
func WithLargura(largura float64) Options {
	return func(r *Params) error {
		if largura <= 0 {
			return errors.New("Largura Zerada")
		}
		r.largura = largura
		return nil
	}
}
func (p Params) GetALtura() float64 {
	return p.altura
}
func (p Params) GetBase() float64 {
	return p.base
}
func (p Params) GetLargura() float64 {
	return p.largura
}
func (p *Params) SetAltura(altura float64) {
	p.altura = altura
}
func (p *Params) SetBase(base float64) {
	p.base = base
}
func (p *Params) SetLargura(largura float64) {
	p.largura = largura
}
