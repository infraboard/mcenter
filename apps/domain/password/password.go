package password

import "github.com/infraboard/mcenter/apps/domain"

const (
	// LengthWeak weak length password
	LengthWeak = 6

	// LengthOK ok length password
	LengthOK = 12

	// LengthStrong strong length password
	LengthStrong = 24

	// LengthVeryStrong very strong length password
	LengthVeryStrong = 36

	// DefaultLetterSet is the letter set that is defaulted to - just the
	// alphabet
	DefaultLetterSet = "abcdefghijklmnopqrstuvwxyz"

	// DefaultLetterAmbiguousSet are letters which are removed from the
	// chosen character set if removing similar characters
	DefaultLetterAmbiguousSet = "ijlo"

	// DefaultNumberSet the default symbol set if character set hasn't been
	// selected
	DefaultNumberSet = "0123456789"

	// DefaultNumberAmbiguousSet are the numbers which are removed from the
	// chosen character set if removing similar characters
	DefaultNumberAmbiguousSet = "01"

	// DefaultSymbolSet the default symbol set if character set hasn't been
	// selected
	DefaultSymbolSet = "!$%^&*()_+{}:@[];'#<>?,./|\\-=?"

	// DefaultSymbolAmbiguousSet are the symbols which are removed from the
	// chosen character set if removing ambiguous characters
	DefaultSymbolAmbiguousSet = "<>[](){}:;'/|\\,"
)

var (
	// DefaultConfig is the default configuration, defaults to:
	//    - length = 24
	//    - Includes symbols, numbers, lowercase and uppercase letters.
	//    - Excludes similar and ambiguous characters
	DefaultConfig = domain.PasswordSecurity{
		Length:             LengthStrong,
		IncludeNumber:      true,
		IncludeLowerLetter: true,
		IncludeUpperLetter: true,
		IncludeSymbols:     true,
		ExcludeSimilar:     true,
		ExcludeAmbiguous:   true,
	}
)
