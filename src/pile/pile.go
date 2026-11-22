package pile

import "fmt"

type Piles struct {
	A          []int
	B          []int
	Operations []string
}

func (p *Piles) Ra() {
	p.A = rotate(p.A)
	p.Operations = append(p.Operations, "ra")
}

func (p *Piles) Rb() {
	p.B = rotate(p.B)
	p.Operations = append(p.Operations, "rb")
}

func (p *Piles) Sa() {
	swap(p.A)
	p.Operations = append(p.Operations, "sa")
}

func (p *Piles) Sb() {
	swap(p.B)
	p.Operations = append(p.Operations, "sb")
}

func (p *Piles) Pb() {
	if len(p.A) == 0 {
		return
	}
	p.B = append([]int{p.A[0]}, p.B...)
	p.A = p.A[1:]
	p.Operations = append(p.Operations, "pb")
}

func (p *Piles) Pa() {
	if len(p.B) == 0 {
		return
	}
	p.A = append([]int{p.B[0]}, p.A...)
	p.B = p.B[1:]
	p.Operations = append(p.Operations, "pa")
}

func (p *Piles) afficher() {
	fmt.Println("a :", p.A, " b :", p.B)
}

func (p *Piles) Rra() {
	p.A = reverseRotate(p.A)
	p.Operations = append(p.Operations, "rra")
}

func (p *Piles) Rrb() {
	p.B = reverseRotate(p.B)
	p.Operations = append(p.Operations, "rrb")
}

func (p *Piles) Ss() {
	swap(p.A)
	swap(p.B)
	p.Operations = append(p.Operations, "ss")
}

func (p *Piles) Rr() {
	p.A = rotate(p.A)
	p.B = rotate(p.B)
	p.Operations = append(p.Operations, "rr")
}

func (p *Piles) Rrr() {
	p.A = reverseRotate(p.A)
	p.B = reverseRotate(p.B)
	p.Operations = append(p.Operations, "rrr")
}

// swap echange les deux premiers elements de s. ne fait rien si s a moins de deux elements.
func swap(s []int) {
	if len(s) < 2 {
		return
	}
	s[0], s[1] = s[1], s[0]
}

// rotate place le premier element de s en dernier position
func rotate(s []int) []int {
	if len(s) < 2 {
		return s
	}
	s = append(s[1:], s[0])
	return s
}

// reverseRotate place le dernier element de s a la premiere position
func reverseRotate(s []int) []int {
	if len(s) < 2 {
		return s
	}
	dernier := s[len(s)-1]
	s = append([]int{dernier}, s[:len(s)-1]...)
	return s
}

// Nouvelles recoit les nombres de l'utilisateur les place dans A et laisse le B vide
func Nouvelles(nums []int) *Piles {
	return &Piles{A: nums, B: []int{}}
}

func (p *Piles) EstTriee() bool {
	if len(p.B) != 0 {
		return false
	}
	for i := 0; i < len(p.A)-1; i++ {
		if p.A[i] > p.A[i+1] {
			return false
		}
	}
	return true
}

// Executer lance l'instruction dont le nom est donné.
// Renvoie false si l'instruction n'existe pas.

func (p *Piles) Executer(nom string) bool {
	switch nom {
	case "sa":
		p.Sa()
	case "sb":
		p.Sb()
	case "ss":
		p.Ss()
	case "pa":
		p.Pa()
	case "pb":
		p.Pb()
	case "rr":
		p.Rr()
	case "ra":
		p.Ra()
	case "rb":
		p.Rb()
	case "rrr":
		p.Rrr()
	case "rra":
		p.Rra()
	case "rrb":
		p.Rrb()
	default:
		return false
	}
	return true
}
