package extenso_test

import (
	"fmt"

	extenso "github.com/savio04/extenso-ptbr"
)

func ExampleInt() {
	fmt.Println(extenso.Int(0))
	fmt.Println(extenso.Int(100))
	fmt.Println(extenso.Int(101))
	fmt.Println(extenso.Int(1000))
	fmt.Println(extenso.Int(1234))
	fmt.Println(extenso.Int(1000000))
	fmt.Println(extenso.Int(-42))
	// Output:
	// zero
	// cem
	// cento e um
	// mil
	// mil duzentos e trinta e quatro
	// um milhão
	// menos quarenta e dois
}

func ExampleFormat_feminine() {
	s, _ := extenso.Format(1, extenso.Options{Gender: extenso.Feminine})
	fmt.Println(s)

	s, _ = extenso.Format(200, extenso.Options{Gender: extenso.Feminine})
	fmt.Println(s)

	s, _ = extenso.Format(21, extenso.Options{Gender: extenso.Feminine})
	fmt.Println(s)
	// Output:
	// uma
	// duzentas
	// vinte e uma
}

func ExampleFloat_brl() {
	s, _ := extenso.Float(1.50, extenso.Options{Currency: "BRL"})
	fmt.Println(s)

	s, _ = extenso.Float(0.01, extenso.Options{Currency: "BRL"})
	fmt.Println(s)

	s, _ = extenso.Float(1000.00, extenso.Options{Currency: "BRL"})
	fmt.Println(s)

	s, _ = extenso.Float(0.50, extenso.Options{Currency: "BRL"})
	fmt.Println(s)
	// Output:
	// um real e cinquenta centavos
	// um centavo
	// mil reais
	// cinquenta centavos
}

func ExampleOrdinal() {
	s, _ := extenso.Ordinal(1, extenso.Masculine)
	fmt.Println(s)

	s, _ = extenso.Ordinal(2, extenso.Feminine)
	fmt.Println(s)

	s, _ = extenso.Ordinal(21, extenso.Masculine)
	fmt.Println(s)

	s, _ = extenso.Ordinal(100, extenso.Feminine)
	fmt.Println(s)
	// Output:
	// primeiro
	// segunda
	// vigésimo primeiro
	// centésima
}
