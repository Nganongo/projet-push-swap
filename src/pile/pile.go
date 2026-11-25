package pile

type Piles struct {
	A          []int
	B          []int
	Operations []string
}

// Ra place le premier element de la pile a en dernier avec la fonction rotate
// puis enregistre ra dans operation
func (p *Piles) Ra() {
	p.A = rotate(p.A)
	p.Operations = append(p.Operations, "ra")
}

// Rb place le premier element de la pile b en dernier avec la fonction rotate
// puis enregistre rb dans operation
func (p *Piles) Rb() {
	p.B = rotate(p.B)
	p.Operations = append(p.Operations, "rb")
}

// Sa echange les deux premier element de la pile A avec la fonction swap
// puis enregistre sa dans operation
func (p *Piles) Sa() {
	swap(p.A)
	p.Operations = append(p.Operations, "sa")
}

// Sb echange les deux premier element de la pile B avec la fonction swap
// puis enregistre sb dans operation
func (p *Piles) Sb() {
	swap(p.B)
	p.Operations = append(p.Operations, "sb")
}

// Pb prend le premier element de la pile A et met dans la pile B
// puis enregistre pb dans Operation
func (p *Piles) Pb() {
	if len(p.A) == 0 {
		return
	}
	p.B = append([]int{p.A[0]}, p.B...)
	p.A = p.A[1:]
	p.Operations = append(p.Operations, "pb")
}

// Pa prend le premier element de la pile B et met dans la pile A
// puis enregistre pa dans Operation
func (p *Piles) Pa() {
	if len(p.B) == 0 {
		return
	}
	p.A = append([]int{p.B[0]}, p.A...)
	p.B = p.B[1:]
	p.Operations = append(p.Operations, "pa")
}

// Rra place le dernier element de la pile A au sommet acev la fonction reversRotate
// pui enregistre rra dans Operation
func (p *Piles) Rra() {
	p.A = reverseRotate(p.A)
	p.Operations = append(p.Operations, "rra")
}

// Rrb place le dernier element de la pile B au sommet acev la fonction reversRotate
// pui enregistre rrb dans Operation
func (p *Piles) Rrb() {
	p.B = reverseRotate(p.B)
	p.Operations = append(p.Operations, "rrb")
}

// Ss echange les deux premiers element des deux pile d'un seule coup
// puis enregistre ss dans Operation
func (p *Piles) Ss() {
	swap(p.A)
	swap(p.B)
	p.Operations = append(p.Operations, "ss")
}

// Rr place les premiers element des deux pile en dernier avec la fonction rotate
// pui enregistre rr dans Operation
func (p *Piles) Rr() {
	p.A = rotate(p.A)
	p.B = rotate(p.B)
	p.Operations = append(p.Operations, "rr")
}

// Rrr place les derneirs element des deux pile au sommet des piles avec la fonction reverseRotate
// puis enregistre rrr dans Operation
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

// EstTriee verifie si les element dans la pile entrée sont triée
// sinon retourne false
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
