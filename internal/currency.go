package internal

import (
	"fmt"
	"math"
	"strings"
)

// BRL converte um valor float64 para extenso em reais (BRL).
func BRL(amount float64) (string, error) {
	if math.IsNaN(amount) || math.IsInf(amount, 0) {
		return "", fmt.Errorf("internal: valor inválido para conversão de moeda")
	}

	neg := amount < 0
	if neg {
		amount = -amount
	}

	totalCents := int64(math.Round(amount * 100))
	reais := totalCents / 100
	cents := totalCents % 100

	var parts []string

	if reais > 0 {
		r, err := Cardinal(reais, Masculine)
		if err != nil {
			return "", err
		}
		unit := "reais"
		if reais == 1 {
			unit = "real"
		}
		// Quando o número termina em milhão/bilhão (ex: "um milhão"), usa-se "de" antes da unidade.
		if endsWithScale(r) {
			unit = "de " + unit
		}
		parts = append(parts, r+" "+unit)
	}

	if cents > 0 {
		c, err := Cardinal(cents, Masculine)
		if err != nil {
			return "", err
		}
		unit := "centavos"
		if cents == 1 {
			unit = "centavo"
		}
		parts = append(parts, c+" "+unit)
	}

	if len(parts) == 0 {
		return "zero reais", nil
	}

	result := strings.Join(parts, " e ")
	if neg {
		result = "menos " + result
	}
	return result, nil
}

var scaleEndings = []string{
	"milhão", "milhões", "bilhão", "bilhões", "trilhão", "trilhões",
	"quatrilhão", "quatrilhões",
}

func endsWithScale(s string) bool {
	for _, e := range scaleEndings {
		if strings.HasSuffix(s, e) {
			return true
		}
	}
	return false
}
