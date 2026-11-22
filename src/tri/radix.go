package tri

import "push-swap/src/pile"

func versRang(nombres []int) []int {
	resultat := make([]int, len(nombres))
	for i, nombre := range nombres {
		plusPetit := 0
		for _, val := range nombres {
			if val < nombre {
				plusPetit++
			}
		}
		resultat[i] = plusPetit
	}
	return resultat
}

func trieRadix(p *pile.Piles) {
	p.A = versRang(p.A)
	n := len(p.A)
	bits := 0
	for (n-1)>>bits != 0 {
		bits++
	}

	for bit := 0; bit < bits; bit++ {
		for j := 0; j < n; j++ {
			if (p.A[0]>>bit)&1 == 0 {
				p.Pb()
			} else {
				p.Ra()
			}
		}
		for len(p.B) > 0 {
			p.Pa()
		}
	}
}
