# Push-swap

Programme en Go qui trie une liste d'entiers à l'aide de deux piles (**a** et **b**) et d'un ensemble limité d'instructions, en cherchant à utiliser le moins d'instructions possible.

Le projet contient deux exécutables :

- **`push-swap`** : reçoit une liste d'entiers et affiche la suite d'instructions qui la trie ;
- **`checker`** : reçoit une liste d'entiers, lit des instructions sur l'entrée standard, les exécute, puis affiche `OK` si la pile est triée ou `KO` sinon.

Le projet n'utilise que la bibliothèque standard de Go.

## Compilation

Depuis la racine du projet :

```bash
go build -o push-swap ./cmd/push-swap
go build -o checker ./cmd/checker
```

Sous Windows, ajouter `.exe` aux noms (`push-swap.exe`, `checker.exe`) et lancer les programmes avec `.\` devant.

## Utilisation

### push-swap

```bash
$ ./push-swap 2 1 3 6 5 8
...
$ ./push-swap "3 2 1"
ra
sa
$ ./push-swap 1 2 3
$ ./push-swap 1 abc 3
Error
$ ./push-swap 1 2 2
Error
```

- Les nombres peuvent être passés séparément (`2 1 3`) ou dans une seule chaîne (`"2 1 3"`).
- Le premier nombre correspond au **sommet** de la pile a.
- Si la liste est déjà triée, ou s'il n'y a aucun argument, rien n'est affiché.
- En cas d'erreur (valeur qui n'est pas un entier, entier hors des limites 32 bits, doublon), le programme écrit `Error` sur la sortie d'erreur.

### checker

```bash
$ ./checker 3 2 1
sa
rra
(Ctrl+D, ou Ctrl+Z puis Entrée sous Windows)
OK
$ ./push-swap 4 1 5 2 3 | ./checker 4 1 5 2 3
OK
$ echo "xx" | ./checker 3 2 1
Error
```

- Les instructions sont lues une par ligne.
- Affiche `OK` si a est triée par ordre croissant et b est vide, sinon `KO`.
- Une instruction inconnue ou mal écrite provoque `Error` sur la sortie d'erreur.

## Instructions disponibles

| Instruction | Effet |
|---|---|
| `sa` / `sb` | échange les deux premiers éléments de a / de b |
| `ss` | `sa` et `sb` en même temps |
| `pa` | prend le sommet de b et le place au sommet de a |
| `pb` | prend le sommet de a et le place au sommet de b |
| `ra` / `rb` | rotation : le premier élément de a / de b devient le dernier |
| `rr` | `ra` et `rb` en même temps |
| `rra` / `rrb` | rotation inverse : le dernier élément de a / de b devient le premier |
| `rrr` | `rra` et `rrb` en même temps |

Une instruction appliquée à une pile vide ou trop petite ne fait rien.

## Organisation du projet

```
push-swap/
├── go.mod
├── cmd/
│   ├── push-swap/main.go   programme push-swap
│   └── checker/main.go     programme checker
└── src/
    ├── lecture/            lecture et validation des arguments
    ├── pile/               les piles a et b et les 11 instructions
    └── tri/                algorithmes de tri
        ├── tri.go          tri des petites listes et choix de l'algorithme
        └── radix.go        tri radix pour les grandes listes
```

Chaque package a une seule responsabilité : `lecture` transforme les arguments en entiers, `pile` gère les piles et les instructions, `tri` produit la suite d'instructions. Les deux programmes réutilisent les mêmes packages.

## Algorithme

Le choix de l'algorithme dépend de la taille de la liste.

### 2 à 5 nombres

- **2 nombres** : un `sa` si nécessaire.
- **3 nombres** : on repère le plus grand. S'il est en haut, `ra` ; s'il est au milieu, `rra`. Le plus grand se retrouve alors en bas, et un `sa` final range les deux premiers si besoin. Au maximum 2 instructions.
- **4 ou 5 nombres** : tant que a contient plus de 3 nombres, on amène le plus petit au sommet par le chemin le plus court (`ra` ou `rra`), puis on le pousse dans b (`pb`). On trie les 3 restants, puis on ramène les nombres de b avec `pa`.

### Plus de 5 nombres : tri radix

1. Chaque nombre est remplacé par son **rang** (sa position dans la liste triée), ce qui donne des valeurs de 0 à n-1. L'ordre relatif ne change pas.
2. Les rangs sont triés **bit par bit**, du bit de poids faible au bit de poids fort. À chaque passage, chaque nombre de a est examiné une fois : si le bit courant vaut 0, il part dans b (`pb`), sinon il reste dans a (`ra`). Ensuite, tous les nombres de b reviennent dans a (`pa`).
3. Après le dernier bit, a est triée.

### Résultats mesurés

| Taille | Nombre d'instructions |
|---|---|
| 3 nombres | 2 au maximum |
| 100 nombres | 1 084 |
| 500 nombres | 6 784 |

Toutes les sorties ont été vérifiées avec `checker` (résultat `OK`).

## Tests rapides

```bash
ARG="4 67 3 87 23"; ./push-swap $ARG | ./checker $ARG
ARG=$(seq 1 100 | shuf | tr '\n' ' '); ./push-swap $ARG | wc -l
```

Sous PowerShell :

```powershell
$n = (1..100 | Get-Random -Count 100) -join ' '
.\push-swap.exe $n | Measure-Object -Line
.\push-swap.exe $n | .\checker.exe $n
```
