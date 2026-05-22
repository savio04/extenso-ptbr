// Package extenso converte números para sua representação por extenso
// em português brasileiro (pt-BR).
//
// Suporta números cardinais, ordinais e valores monetários em BRL.
// Zero dependências externas.
//
// Uso básico:
//
//	extenso.Int(1234)
//	// => "mil duzentos e trinta e quatro"
//
//	extenso.Ordinal(2, extenso.Masculine)
//	// => "segundo"
//
//	extenso.Float(1.50, extenso.Options{Currency: "BRL"})
//	// => "um real e cinquenta centavos", nil
package extenso

import (
	"fmt"

	"github.com/savio04/extenso-ptbr/internal"
)

// Gender define o gênero gramatical para concordância nominal.
type Gender int

const (
	// Masculine é o gênero masculino (padrão).
	Masculine Gender = iota
	// Feminine é o gênero feminino.
	Feminine
)

// Options configura o comportamento da conversão.
type Options struct {
	// Gender controla a concordância de gênero (padrão: Masculine).
	Gender Gender
	// Currency, quando igual a "BRL", ativa o modo de moeda brasileira.
	// Neste modo Float retorna valores em reais e centavos.
	Currency string
}

// Int converte n para sua representação por extenso em pt-BR,
// usando gênero masculino. Em pt-BR "1000" é sempre "mil", nunca "um mil".
//
//	Int(0)    // => "zero"
//	Int(100)  // => "cem"
//	Int(101)  // => "cento e um"
//	Int(1000) // => "mil"
//	Int(-5)   // => "menos cinco"
func Int(n int64) string {
	s, _ := internal.Cardinal(n, internal.Masculine)
	return s
}

// Format converte n para extenso com controle de gênero e moeda.
// Retorna erro se Currency for inválida (somente "BRL" ou "" são aceitos).
//
//	Format(42, Options{Gender: Feminine})
//	// => "quarenta e duas", nil
func Format(n int64, opts Options) (string, error) {
	if opts.Currency != "" && opts.Currency != "BRL" {
		return "", fmt.Errorf("extenso: moeda não suportada %q (use \"BRL\" ou \"\")", opts.Currency)
	}
	g := toInternalGender(opts.Gender)
	s, err := internal.Cardinal(n, g)
	if err != nil {
		return "", fmt.Errorf("extenso: %w", err)
	}
	return s, nil
}

// Float converte um valor decimal para extenso.
// Com Currency "BRL", formata como valor monetário em reais e centavos.
//
//	Float(1.50, Options{Currency: "BRL"})
//	// => "um real e cinquenta centavos", nil
//
//	Float(0.01, Options{Currency: "BRL"})
//	// => "um centavo", nil
func Float(n float64, opts Options) (string, error) {
	if opts.Currency == "BRL" {
		s, err := internal.BRL(n)
		if err != nil {
			return "", fmt.Errorf("extenso: %w", err)
		}
		return s, nil
	}
	if opts.Currency != "" {
		return "", fmt.Errorf("extenso: moeda não suportada %q (use \"BRL\" ou \"\")", opts.Currency)
	}
	g := toInternalGender(opts.Gender)
	s, err := internal.Cardinal(int64(n), g)
	if err != nil {
		return "", fmt.Errorf("extenso: %w", err)
	}
	return s, nil
}

// Ordinal converte n para sua forma ordinal em pt-BR.
// Suporta valores de 1 a 9999.
//
//	Ordinal(1, Masculine)  // => "primeiro", nil
//	Ordinal(2, Feminine)   // => "segunda", nil
//	Ordinal(21, Masculine) // => "vigésimo primeiro", nil
func Ordinal(n int64, g Gender) (string, error) {
	s, err := internal.Ordinal(n, toInternalGender(g))
	if err != nil {
		return "", fmt.Errorf("extenso: %w", err)
	}
	return s, nil
}

func toInternalGender(g Gender) internal.Gender {
	if g == Feminine {
		return internal.Feminine
	}
	return internal.Masculine
}
