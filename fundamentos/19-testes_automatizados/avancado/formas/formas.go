package formas

import (
	"math"
)

type Retangulo struct {
	altura  float64
	largura float64
}
type Circulo struct {
	raio float64
}

type Forma interface {
	area() float64
}

// cria um método area associando a struct do retangulo
func (r Retangulo) Area() float64 {
	return r.altura * r.largura
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}
