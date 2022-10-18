package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10.0, 12.0}
		areaEsperada := 120.0
		areaRecebida := ret.Area()

		if areaRecebida != areaEsperada {
			t.Fatalf("A área recebida foi %f, mas a área esperada era %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := math.Pi * 100
		areaRecebida := circ.Area()

		if areaRecebida != areaEsperada {
			t.Fatalf("A área recebida foi %f, mas a área esperada era %f", areaRecebida, areaEsperada)
		}
	})
}
