package internal

import (
	"fmt"
	"strings"
)

var ordinalOnesM = []string{
	"", "primeiro", "segundo", "terceiro", "quarto", "quinto",
	"sexto", "sétimo", "oitavo", "nono", "décimo",
}

var ordinalOnesF = []string{
	"", "primeira", "segunda", "terceira", "quarta", "quinta",
	"sexta", "sétima", "oitava", "nona", "décima",
}

// Ordinal converte n (1–9999) para ordinal em pt-BR.
func Ordinal(n int64, g Gender) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("internal: ordinal requer número positivo, recebeu %d", n)
	}
	if n > 9999 {
		return "", fmt.Errorf("internal: ordinal suportado até 9999, recebeu %d", n)
	}
	if n <= 10 {
		if g == Feminine {
			return ordinalOnesF[n], nil
		}
		return ordinalOnesM[n], nil
	}
	return buildOrdinal(n, g), nil
}

func buildOrdinal(n int64, g Gender) string {
	var parts []string

	mil := n / 1000
	rem := n % 1000

	if mil > 0 {
		if mil == 1 {
			parts = append(parts, genderSuffix("milésimo", g))
		} else {
			cardinal, _ := Cardinal(mil, Masculine)
			parts = append(parts, cardinal+" "+genderSuffix("milésimo", g))
		}
	}

	cent := rem / 100
	rem2 := rem % 100

	if cent > 0 {
		parts = append(parts, genderSuffix(centOrdinals[cent], g))
	}
	if rem2 > 0 {
		parts = append(parts, smallOrdinal(rem2, g))
	}

	return strings.Join(parts, " ")
}

var centOrdinals = [10]string{
	"", "centésimo", "ducentésimo", "trecentésimo", "quadringentésimo",
	"quingentésimo", "seiscentésimo", "septingentésimo", "octingentésimo", "noningentésimo",
}

var tensOrdinals = [10]string{
	"", "décimo", "vigésimo", "trigésimo", "quadragésimo", "quinquagésimo",
	"sexagésimo", "septuagésimo", "octogésimo", "nonagésimo",
}

// onesOrdM e onesOrdF cobrem 1–19, incluindo as formas compostas 11–19.
var onesOrdM = [20]string{
	"", "primeiro", "segundo", "terceiro", "quarto", "quinto",
	"sexto", "sétimo", "oitavo", "nono", "décimo",
	"décimo primeiro", "décimo segundo", "décimo terceiro", "décimo quarto",
	"décimo quinto", "décimo sexto", "décimo sétimo", "décimo oitavo", "décimo nono",
}

var onesOrdF = [20]string{
	"", "primeira", "segunda", "terceira", "quarta", "quinta",
	"sexta", "sétima", "oitava", "nona", "décima",
	"décima primeira", "décima segunda", "décima terceira", "décima quarta",
	"décima quinta", "décima sexta", "décima sétima", "décima oitava", "décima nona",
}

func smallOrdinal(n int64, g Gender) string {
	if n < 20 {
		if g == Feminine {
			return onesOrdF[n]
		}
		return onesOrdM[n]
	}
	t := genderSuffix(tensOrdinals[n/10], g)
	u := n % 10
	if u == 0 {
		return t
	}
	var unit string
	if g == Feminine {
		unit = onesOrdF[u]
	} else {
		unit = onesOrdM[u]
	}
	return t + " " + unit
}

// genderSuffix troca o sufixo "o" final por "a" para gênero feminino.
func genderSuffix(s string, g Gender) string {
	if g == Masculine || s == "" {
		return s
	}
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == 'o' {
			runes[i] = 'a'
			return string(runes)
		}
	}
	return s
}
