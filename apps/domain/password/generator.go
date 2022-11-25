package password

import (
	"crypto/rand"
	"math/big"
	"strings"

	"github.com/infraboard/mcenter/apps/domain"
)

// Generator is what generates the password
type Generator struct {
	*domain.PasswordSecurity
}

// New returns a new generator
func New(config *domain.PasswordSecurity) *Generator {
	if config == nil {
		config = &DefaultConfig
	}

	if !config.IncludeSymbols &&
		!config.IncludeUpperLetter &&
		!config.IncludeLowerLetter &&
		!config.IncludeNumber &&
		config.CharacterSet == "" {
		config = &DefaultConfig
	}

	if config.Length == 0 {
		config.Length = LengthStrong
	}

	if config.CharacterSet == "" {
		config.CharacterSet = BuildCharacterSet(config)
	}

	return &Generator{PasswordSecurity: config}
}

func BuildCharacterSet(config *domain.PasswordSecurity) string {
	var characterSet string
	if config.IncludeLowerLetter {
		characterSet += DefaultLetterSet
		if config.ExcludeSimilar {
			characterSet = removeCharacters(characterSet, DefaultLetterAmbiguousSet)
		}
	}

	if config.IncludeUpperLetter {
		characterSet += strings.ToUpper(DefaultLetterSet)
		if config.ExcludeSimilar {
			characterSet = removeCharacters(characterSet, strings.ToUpper(DefaultLetterAmbiguousSet))
		}
	}

	if config.IncludeNumber {
		characterSet += DefaultNumberSet
		if config.ExcludeSimilar {
			characterSet = removeCharacters(characterSet, DefaultNumberAmbiguousSet)
		}
	}

	if config.IncludeSymbols {
		characterSet += DefaultSymbolSet
		if config.ExcludeAmbiguous {
			characterSet = removeCharacters(characterSet, DefaultSymbolAmbiguousSet)
		}
	}

	return characterSet
}

func removeCharacters(str, characters string) string {
	return strings.Map(func(r rune) rune {
		if !strings.ContainsRune(characters, r) {
			return r
		}
		return -1
	}, str)
}

// NewWithDefault returns a new generator with the default
// config
func NewWithDefault() *Generator {
	return New(&DefaultConfig)
}

// Generate generates one password with length set in the
// config
func (g Generator) Generate() (*string, error) {
	var generated string
	characterSet := strings.Split(g.PasswordSecurity.CharacterSet, "")
	max := big.NewInt(int64(len(characterSet)))

	for i := 0; i < int(g.PasswordSecurity.Length); i++ {
		val, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, err
		}
		generated += characterSet[val.Int64()]
	}
	return &generated, nil
}

// GenerateMany generates multiple passwords with length set
// in the config
func (g Generator) GenerateMany(amount int) ([]string, error) {
	var generated []string
	for i := 0; i < amount; i++ {
		str, err := g.Generate()
		if err != nil {
			return nil, err
		}

		generated = append(generated, *str)
	}
	return generated, nil
}

// GenerateWithLength generate one password with set length
func (g Generator) GenerateWithLength(length int) (*string, error) {
	var generated string
	characterSet := strings.Split(g.PasswordSecurity.CharacterSet, "")
	max := big.NewInt(int64(len(characterSet)))
	for i := 0; i < length; i++ {
		val, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, err
		}
		generated += characterSet[val.Int64()]
	}
	return &generated, nil
}

// GenerateManyWithLength generates multiple passwords with set length
func (g Generator) GenerateManyWithLength(amount, length int) ([]string, error) {
	var generated []string
	for i := 0; i < amount; i++ {
		str, err := g.GenerateWithLength(length)
		if err != nil {
			return nil, err
		}
		generated = append(generated, *str)
	}
	return generated, nil
}
