package formas

import (
	"math"
)

type Forma interface {
	Area() float64
}

type Retangulo struct {
	altura  float64
	largura float64
}

func (r Retangulo) Area() float64 {
	return r.altura * r.largura
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

type Circulo struct {
	raio float64
}
