package tri

import "push-swap/src/pile"

func trierTrois(p *pile.Piles) {
	if len(p.A) != 3 {
		return
	}
	if p.A[0] > p.A[1] && p.A[0] > p.A[2] {
		p.Ra()
	} else if p.A[1] > p.A[0] && p.A[1] > p.A[2] {
		p.Rra()
	}
	if p.A[0] > p.A[1] {
		p.Sa()
	}
}

func positionDuMinimum(nbre []int) int {
	indexMin := 0
	for i := 1; i < len(nbre); i++ {
		if nbre[i] < nbre[indexMin] {
			indexMin = i
		}
	}
	return indexMin
}

func amenerAuSommet(p *pile.Piles, position int) {
	if position <= len(p.A)/2 {
		for i := 0; i < position; i++ {
			p.Ra()
		}
	} else {
		for i := 0; i < len(p.A)-position; i++ {
			p.Rra()
		}
	}
}

func triPetit(p *pile.Piles) {
	if len(p.A) == 2 {
		if p.A[0] > p.A[1] {
			p.Sa()
			return
		}
	}
	for len(p.A) > 3 {
		position := positionDuMinimum(p.A)
		amenerAuSommet(p, position)
		p.Pb()
	}
	trierTrois(p)
	for len(p.B) > 0 {
		p.Pa()
	}
}

// Trier trie la pile A à l'aide de B. Si A est déjà triée, aucune instruction n'est ajoutée.
func Trier(p *pile.Piles) {
	if p.EstTriee() {
		return
	}
	if len(p.A) <= 5 {
		triPetit(p)
	} else {
		trieRadix(p)
	}
}
