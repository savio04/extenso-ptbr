package internal

import (
	"fmt"
	"strings"
)

// Gender define o gênero gramatical.
type Gender int

const (
	Masculine Gender = iota
	Feminine
)

var ones = [20]string{
	"zero", "um", "dois", "três", "quatro", "cinco",
	"seis", "sete", "oito", "nove", "dez", "onze",
	"doze", "treze", "quatorze", "quinze", "dezesseis",
	"dezessete", "dezoito", "dezenove",
}

var onesF = [20]string{
	"zero", "uma", "duas", "três", "quatro", "cinco",
	"seis", "sete", "oito", "nove", "dez", "onze",
	"doze", "treze", "quatorze", "quinze", "dezesseis",
	"dezessete", "dezoito", "dezenove",
}

var tens = [10]string{
	"", "dez", "vinte", "trinta", "quarenta", "cinquenta",
	"sessenta", "setenta", "oitenta", "noventa",
}

var hundreds = [10]string{
	"", "cem", "duzentos", "trezentos", "quatrocentos", "quinhentos",
	"seiscentos", "setecentos", "oitocentos", "novecentos",
}

var hundredsF = [10]string{
	"", "cem", "duzentas", "trezentas", "quatrocentas", "quinhentas",
	"seiscentas", "setecentas", "oitocentas", "novecentas",
}

type scaleInfo struct {
	singular string
	plural   string
}

var scales = []scaleInfo{
	{"", ""},
	{"mil", "mil"},
	{"milhão", "milhões"},
	{"bilhão", "bilhões"},
	{"trilhão", "trilhões"},
	{"quatrilhão", "quatrilhões"},
}

// Cardinal converte n para extenso cardinal em pt-BR.
func Cardinal(n int64, g Gender) (string, error) {
	if n == 0 {
		return "zero", nil
	}

	neg := n < 0
	if n == -1<<63 {
		return "", fmt.Errorf("internal: valor fora do intervalo suportado")
	}
	if neg {
		n = -n
	}

	groups := splitGroups(n)
	if len(groups) > len(scales) {
		return "", fmt.Errorf("internal: número muito grande")
	}

	type seg struct {
		text string
		val  int64 // valor do grupo (1–999), usado para regra do "e"
	}
	var segs []seg

	for i := len(groups) - 1; i >= 0; i-- {
		v := groups[i]
		if v == 0 {
			continue
		}

		var word string
		// "mil" nunca leva "um" na frente — apenas "dois mil", "três mil", etc.
		if i == 1 && v == 1 {
			word = "mil"
		} else {
			// grupos acima das unidades são sempre masculinos
			gg := Masculine
			if i == 0 {
				gg = g
			}
			word = groupToWords(v, gg)
			if i > 0 {
				sc := scales[i]
				name := sc.singular
				if v > 1 {
					name = sc.plural
				}
				word += " " + name
			}
		}
		segs = append(segs, seg{word, v})
	}

	if len(segs) == 0 {
		return "zero", nil
	}

	// Constrói o resultado, inserindo "e" quando o grupo seguinte é < 100
	// ou é múltiplo exato de 100 (ex: 100, 200, 500…).
	result := segs[0].text
	for i := 1; i < len(segs); i++ {
		v := segs[i].val
		if v < 100 || v%100 == 0 {
			result += " e " + segs[i].text
		} else {
			result += " " + segs[i].text
		}
	}

	if neg {
		result = "menos " + result
	}
	return result, nil
}

// splitGroups divide n em grupos de 3 dígitos; índice 0 = unidades.
func splitGroups(n int64) []int64 {
	var groups []int64
	for n > 0 {
		groups = append(groups, n%1000)
		n /= 1000
	}
	return groups
}

// groupToWords converte 1–999 para palavras.
func groupToWords(n int64, g Gender) string {
	h := n / 100
	rem := n % 100

	var parts []string
	if h > 0 {
		parts = append(parts, hundredsWord(h, rem > 0, g))
	}
	if rem > 0 {
		parts = append(parts, remToWords(rem, g))
	}
	return strings.Join(parts, " e ")
}

func hundredsWord(h int64, hasRem bool, g Gender) string {
	if h == 1 && hasRem {
		return "cento"
	}
	if g == Feminine {
		return hundredsF[h]
	}
	return hundreds[h]
}

func remToWords(rem int64, g Gender) string {
	if rem < 20 {
		if g == Feminine {
			return onesF[rem]
		}
		return ones[rem]
	}
	t := tens[rem/10]
	u := rem % 10
	if u == 0 {
		return t
	}
	unit := ones[u]
	if g == Feminine {
		unit = onesF[u]
	}
	return t + " e " + unit
}
