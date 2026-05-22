package extenso_test

import (
	"testing"

	extenso "github.com/savio04/extenso-ptbr"
)

func TestInt(t *testing.T) {
	cases := []struct {
		n    int64
		want string
	}{
		{0, "zero"},
		{1, "um"},
		{2, "dois"},
		{10, "dez"},
		{11, "onze"},
		{15, "quinze"},
		{19, "dezenove"},
		{20, "vinte"},
		{21, "vinte e um"},
		{99, "noventa e nove"},
		{100, "cem"},
		{101, "cento e um"},
		{200, "duzentos"},
		{315, "trezentos e quinze"},
		{999, "novecentos e noventa e nove"},
		{1000, "mil"},
		{1001, "mil e um"},
		{1100, "mil e cem"},
		{1500, "mil e quinhentos"},
		{1234, "mil duzentos e trinta e quatro"},
		{2000, "dois mil"},
		{10000, "dez mil"},
		{100000, "cem mil"},
		{999999, "novecentos e noventa e nove mil novecentos e noventa e nove"},
		{1000000, "um milhão"},
		{2000000, "dois milhões"},
		{1500000, "um milhão e quinhentos mil"},
		{1000000000, "um bilhão"},
		{2000000000, "dois bilhões"},
		{-1, "menos um"},
		{-1000, "menos mil"},
		{-1234567, "menos um milhão duzentos e trinta e quatro mil quinhentos e sessenta e sete"},
	}

	for _, tc := range cases {
		t.Run(tc.want, func(t *testing.T) {
			got := extenso.Int(tc.n)
			if got != tc.want {
				t.Errorf("Int(%d) = %q, want %q", tc.n, got, tc.want)
			}
		})
	}
}

func TestFormat_Feminine(t *testing.T) {
	cases := []struct {
		n    int64
		want string
	}{
		{1, "uma"},
		{2, "duas"},
		{21, "vinte e uma"},
		{200, "duzentas"},
		{201, "duzentas e uma"},
		{1001, "mil e uma"},
		{2000, "dois mil"},
		{1000001, "um milhão e uma"},
	}

	for _, tc := range cases {
		t.Run(tc.want, func(t *testing.T) {
			got, err := extenso.Format(tc.n, extenso.Options{Gender: extenso.Feminine})
			if err != nil {
				t.Fatalf("Format(%d, Feminine) unexpected error: %v", tc.n, err)
			}
			if got != tc.want {
				t.Errorf("Format(%d, Feminine) = %q, want %q", tc.n, got, tc.want)
			}
		})
	}
}

func TestFormat_InvalidCurrency(t *testing.T) {
	_, err := extenso.Format(10, extenso.Options{Currency: "USD"})
	if err == nil {
		t.Error("expected error for unsupported currency, got nil")
	}
}

func TestFloat_BRL(t *testing.T) {
	cases := []struct {
		amount float64
		want   string
	}{
		{0.0, "zero reais"},
		{0.01, "um centavo"},
		{0.50, "cinquenta centavos"},
		{0.99, "noventa e nove centavos"},
		{1.00, "um real"},
		{1.01, "um real e um centavo"},
		{1.50, "um real e cinquenta centavos"},
		{2.00, "dois reais"},
		{2.99, "dois reais e noventa e nove centavos"},
		{1000.00, "mil reais"},
		{1000.01, "mil reais e um centavo"},
		{1000000.00, "um milhão de reais"},
		{-1.50, "menos um real e cinquenta centavos"},
	}

	for _, tc := range cases {
		t.Run(tc.want, func(t *testing.T) {
			got, err := extenso.Float(tc.amount, extenso.Options{Currency: "BRL"})
			if err != nil {
				t.Fatalf("Float(%v, BRL) unexpected error: %v", tc.amount, err)
			}
			if got != tc.want {
				t.Errorf("Float(%v, BRL) = %q, want %q", tc.amount, got, tc.want)
			}
		})
	}
}

func TestOrdinal(t *testing.T) {
	cases := []struct {
		n    int64
		g    extenso.Gender
		want string
	}{
		{1, extenso.Masculine, "primeiro"},
		{2, extenso.Masculine, "segundo"},
		{3, extenso.Masculine, "terceiro"},
		{4, extenso.Masculine, "quarto"},
		{5, extenso.Masculine, "quinto"},
		{6, extenso.Masculine, "sexto"},
		{7, extenso.Masculine, "sétimo"},
		{8, extenso.Masculine, "oitavo"},
		{9, extenso.Masculine, "nono"},
		{10, extenso.Masculine, "décimo"},
		{11, extenso.Masculine, "décimo primeiro"},
		{20, extenso.Masculine, "vigésimo"},
		{21, extenso.Masculine, "vigésimo primeiro"},
		{100, extenso.Masculine, "centésimo"},
		{1000, extenso.Masculine, "milésimo"},
		{1, extenso.Feminine, "primeira"},
		{2, extenso.Feminine, "segunda"},
		{10, extenso.Feminine, "décima"},
		{11, extenso.Feminine, "décima primeira"},
		{21, extenso.Feminine, "vigésima primeira"},
		{100, extenso.Feminine, "centésima"},
	}

	for _, tc := range cases {
		t.Run(tc.want, func(t *testing.T) {
			got, err := extenso.Ordinal(tc.n, tc.g)
			if err != nil {
				t.Fatalf("Ordinal(%d, %v) unexpected error: %v", tc.n, tc.g, err)
			}
			if got != tc.want {
				t.Errorf("Ordinal(%d, %v) = %q, want %q", tc.n, tc.g, got, tc.want)
			}
		})
	}
}

func TestOrdinal_Errors(t *testing.T) {
	_, err := extenso.Ordinal(0, extenso.Masculine)
	if err == nil {
		t.Error("expected error for Ordinal(0), got nil")
	}
	_, err = extenso.Ordinal(-1, extenso.Masculine)
	if err == nil {
		t.Error("expected error for Ordinal(-1), got nil")
	}
	_, err = extenso.Ordinal(10000, extenso.Masculine)
	if err == nil {
		t.Error("expected error for Ordinal(10000), got nil")
	}
}
