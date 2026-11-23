package lecture

import (
	"fmt"
	"strconv"
	"strings"
)

func lireNombres(args []string) ([]int, error) {
	var nombres []int
	for _, arg := range args {
		morceaux := strings.Fields(arg)
		for _, morceau := range morceaux {
			nombre, err := strconv.ParseInt(morceau, 10, 32)
			if err != nil {
				return nil, fmt.Errorf("%q n'est pas un nombre entier", morceau)
			}

			nombres = append(nombres, int(nombre))
		}
	}
	return nombres, nil
}

// VerifieDoublons
func verifierDoublons(nombres []int) error {
	vue := map[int]bool{}
	for _, nombre := range nombres {
		if vue[nombre] {
			return fmt.Errorf("le nombre %d apparait plusieur fois", nombre)
		}
		vue[nombre] = true
	}
	return nil
}

// Lire convertit les arguments en entiers et vérifie qu'il n'y a pas de doublons. Renvoie une erreur si un argument n'est pas un entier ou s'il y a un doublon.
func Lire(args []string) ([]int, error) {
	nombres, err := lireNombres(args)
	if err != nil {
		return nil, err
	}
	err = verifierDoublons(nombres)
	if err != nil {
		return nil, err
	}
	return nombres, nil
}
