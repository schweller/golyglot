# golyglot

**Please note**: this is an unoficial repository and not mantained by DeepL. It contains an unoficial CLI tool and an API client for Go.

---

golyglot is a toolbelt for the [DeepL API](https://www.deepl.com/docs-api), either via a Go client library or CLI.

## Install

### Go client

`go get github.com/schweller/golyglot/api`

## Usage

```go

// Generate a config
config := golyglot.DefaultConfig()

// DeepL client
client := golyglot.NewClient(config)

// Instance a list of sentences
var sentences []string

sentences = append(sentence, "Hello world!")

resp, err := client.Translate().PostTranslations(sentences, "IT")
```

## Roadmap

- Support document upload
- Support glossary management
- Proper documentation
- CLI output:
    - Distinguish layers
    - Pretty print