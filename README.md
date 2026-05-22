# extenso-ptbr

[![CI](https://github.com/savio04/extenso-ptbr/actions/workflows/ci.yml/badge.svg)](https://github.com/savio04/extenso-ptbr/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/savio04/extenso-ptbr.svg)](https://pkg.go.dev/github.com/savio04/extenso-ptbr)

Converte números para sua representação por extenso em **português brasileiro (pt-BR)**.  
Ideal para sistemas financeiros, jurídicos e emissão de documentos no Brasil.

---

## Instalação

```bash
go get github.com/savio04/extenso-ptbr
```

Requer Go 1.26+. Zero dependências externas.

---

## Uso

```go
import extenso "github.com/savio04/extenso-ptbr"

// Cardinal básico
extenso.Int(0)        // "zero"
extenso.Int(100)      // "cem"
extenso.Int(101)      // "cento e um"
extenso.Int(1000)     // "mil"
extenso.Int(1500)     // "mil e quinhentos"
extenso.Int(1234)     // "mil duzentos e trinta e quatro"
extenso.Int(1000000)  // "um milhão"
extenso.Int(-42)      // "menos quarenta e dois"

// Gênero feminino
extenso.Format(1, extenso.Options{Gender: extenso.Feminine})
// => "uma", nil

extenso.Format(200, extenso.Options{Gender: extenso.Feminine})
// => "duzentas", nil

// Moeda BRL
extenso.Float(1.50, extenso.Options{Currency: "BRL"})
// => "um real e cinquenta centavos", nil

extenso.Float(1000000.00, extenso.Options{Currency: "BRL"})
// => "um milhão de reais", nil

// Ordinais
extenso.Ordinal(1, extenso.Masculine)  // => "primeiro", nil
extenso.Ordinal(2, extenso.Feminine)   // => "segunda", nil
extenso.Ordinal(21, extenso.Masculine) // => "vigésimo primeiro", nil
```

---

## Funcionalidades

| Feature                        | Suportado |
|-------------------------------|-----------|
| Cardinais (`int64`)           | ✅        |
| Gênero (masculino/feminino)   | ✅        |
| Ordinais (1–9999)             | ✅        |
| Moeda BRL (`float64`)         | ✅        |
| Números negativos             | ✅        |
| Zero dependências externas    | ✅        |
| Thread-safe                   | ✅        |

---

## Edge cases / Regras pt-BR

| Caso                         | Resultado                              | Regra aplicada                                    |
|-----------------------------|----------------------------------------|---------------------------------------------------|
| `Int(100)`                  | `"cem"`                                | "cem" sozinho, "cento" se houver mais             |
| `Int(101)`                  | `"cento e um"`                         | "cento" quando seguido de outros algarismos       |
| `Int(1000)`                 | `"mil"`                                | Nunca "um mil" — regra gramatical pt-BR           |
| `Int(1100)`                 | `"mil e cem"`                          | "e" quando grupo seguinte é múltiplo de 100       |
| `Int(1234)`                 | `"mil duzentos e trinta e quatro"`     | Sem "e" entre grupos quando ≥100 e não múltiplo   |
| `Int(1000000)`              | `"um milhão"`                          | Milhão/bilhão levam "um"                          |
| `Float(1M, BRL)`            | `"um milhão de reais"`                 | "de" antes da unidade monetária após escala       |
| `Format(200, Feminine)`     | `"duzentas"`                           | Centenas têm forma feminina própria               |

---

## API

```go
// Int converte int64 para extenso (masculino).
func Int(n int64) string

// Format converte com controle de gênero e moeda.
func Format(n int64, opts Options) (string, error)

// Float converte float64; com Currency "BRL" formata como moeda.
func Float(n float64, opts Options) (string, error)

// Ordinal converte para forma ordinal (1–9999).
func Ordinal(n int64, g Gender) (string, error)

type Gender int
const (
    Masculine Gender = iota
    Feminine
)

type Options struct {
    Gender   Gender
    Currency string // "" ou "BRL"
}
```

---

---

# extenso-ptbr (English)

Converts numbers to their written form in **Brazilian Portuguese (pt-BR)**.  
Useful for financial, legal, and document-generation systems.

## Installation

```bash
go get github.com/savio04/extenso-ptbr
```

Requires Go 1.26+. Zero external dependencies.

## Quick start

```go
import extenso "github.com/savio04/extenso-ptbr"

extenso.Int(1000)   // "mil"
extenso.Int(1234)   // "mil duzentos e trinta e quatro"

extenso.Format(2, extenso.Options{Gender: extenso.Feminine})
// => "duas", nil

extenso.Float(1.50, extenso.Options{Currency: "BRL"})
// => "um real e cinquenta centavos", nil

extenso.Ordinal(21, extenso.Masculine) // => "vigésimo primeiro", nil
```

## License

MIT © savio04
