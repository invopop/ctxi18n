# ctxi18n

Go Context Internationalization - translating apps easily.

## Introduction

`ctxi18n` is heavily influenced by the Ruby i18n gem and aims to make it internationalization in Go applications as straightforward as in Ruby.

As the name suggests, `ctxi18n` focusses on making i18n methods accessible via the application's context for a given incoming request. I18n should be as quick and easy to use as possible, so this package provides a set of helper methods with short names and parameter lists.

Key Features:

- Loads locale files in YAML, like Ruby i18n.
- Supports Go's embed FS to load data from inside binary.
- Short method names and usage to Gettext like `i18n.T()` or `i18n.N()`.
- Support for simple interpolation using keys, e.g. `Some %{key} text`
- Built in support for [templ templating](https://templ.guide/) which uses context throughout.

## Usage

```go

import (
    "github.com/invopop/ctxi18n"
    "github.com/invopop/ctxi18n/i18n"
)

func main() {
    // Prepare translations sources in package global
    ctxi18n.Load(assets.Locale)

    // Load HTTP

    // Add locale to Context
    ctx := ctx18n.WithLocale(request.Context(), "en")

    fmt.Printf(i18n.T(ctx, "hello"))

}
```
