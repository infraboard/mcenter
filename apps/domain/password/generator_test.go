package password_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/domain/password"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		config *domain.PasswordSecurity
	}
	tests := []struct {
		name    string
		args    args
		want    *password.Generator
		wantErr error
	}{
		{
			name: "default config",
			args: args{nil},
			want: func() *password.Generator {
				cfg := &password.DefaultConfig
				cfg.CharacterSet = password.BuildCharacterSet(cfg)
				cfg.Length = password.LengthStrong
				return &password.Generator{cfg}
			}(),
		},
		{
			name: "set config",
			args: args{&domain.PasswordSecurity{
				IncludeLowerLetter: true,
			}},
			want: func() *password.Generator {
				cfg := domain.PasswordSecurity{IncludeLowerLetter: true}
				cfg.CharacterSet = "abcdefghijklmnopqrstuvwxyz"
				cfg.Length = password.LengthStrong
				return &password.Generator{&cfg}
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := password.New(tt.args.config)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWithDefault(t *testing.T) {
	tests := []struct {
		name    string
		want    *password.Generator
		wantErr error
	}{
		{
			name: "default config",
			want: func() *password.Generator {
				cfg := &password.DefaultConfig
				cfg.CharacterSet = password.BuildCharacterSet(cfg)
				return &password.Generator{cfg}
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := password.NewWithDefault()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildCharacterSet(t *testing.T) {
	type args struct {
		config *domain.PasswordSecurity
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "exclude similar characters",
			args: args{
				config: &domain.PasswordSecurity{
					IncludeLowerLetter: true,
					IncludeSymbols:     true,
					IncludeNumber:      true,
					ExcludeSimilar:     true,
					ExcludeAmbiguous:   true,
				},
			},
			want: "abcdefghkmnpqrstuvwxyz23456789!$%^&*_+@#?.-=?",
		},
		{
			name: "exclude numbers",
			args: args{
				config: &domain.PasswordSecurity{
					IncludeLowerLetter: true,
					IncludeSymbols:     true,
					IncludeNumber:      false,
					ExcludeSimilar:     true,
					ExcludeAmbiguous:   true,
				},
			},
			want: "abcdefghkmnpqrstuvwxyz!$%^&*_+@#?.-=?",
		},
		{
			name: "full list",
			args: args{
				config: &domain.PasswordSecurity{
					IncludeLowerLetter: true,
					IncludeSymbols:     true,
					IncludeNumber:      true,
					ExcludeSimilar:     false,
					ExcludeAmbiguous:   false,
				},
			},
			want: "abcdefghijklmnopqrstuvwxyz0123456789!$%^&*()_+{}:@[];'#<>?,./|\\-=?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := password.BuildCharacterSet(tt.args.config); got != tt.want {
				t.Errorf("buildCharacterSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerator_Generate(t *testing.T) {
	type fields struct {
		Config *domain.PasswordSecurity
	}
	tests := []struct {
		name    string
		fields  fields
		test    func(*string, string)
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{&password.DefaultConfig},
			test: func(pwd *string, characterSet string) {
				assert.Len(t, *pwd, int(password.DefaultConfig.Length))
				err := stringMatchesCharacters(*pwd, characterSet)
				if err != nil {
					t.Errorf("Generate() error = %v", err)
					return
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := password.New(tt.fields.Config)
			got, err := g.Generate()
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.test(got, g.CharacterSet)
		})
	}
}

func TestGenerator_GenerateMany(t *testing.T) {
	type fields struct {
		Config *domain.PasswordSecurity
	}
	type args struct {
		amount int
	}
	tests := []struct {
		name    string
		args    args
		fields  fields
		test    func([]string, string)
		wantErr error
	}{
		{
			name:   "valid",
			args:   args{amount: 5},
			fields: fields{Config: &password.DefaultConfig},
			test: func(pwds []string, characterSet string) {
				assert.Len(t, pwds, 5)

				for _, pwd := range pwds {
					err := stringMatchesCharacters(pwd, characterSet)
					if err != nil {
						t.Errorf("Generate() error = %v", err)
						return
					}
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := password.New(tt.fields.Config)
			got, err := g.GenerateMany(tt.args.amount)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GenerateMany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.test(got, g.CharacterSet)
		})
	}
}

func stringMatchesCharacters(str, characters string) error {
	set := strings.Split(characters, "")
	strSet := strings.Split(str, "")

	for _, strChr := range strSet {
		found := false
		for _, setChr := range set {
			if strChr == setChr {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("%v should not be in the str", strChr)
		}
	}

	return nil
}

func TestGenerator_GenerateWithLength(t *testing.T) {
	type fields struct {
		Config *domain.PasswordSecurity
	}
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		test    func(*string, string)
		wantErr error
	}{
		{
			name:   "valid",
			fields: fields{&password.DefaultConfig},
			args:   args{length: 5},
			test: func(pwd *string, characterSet string) {
				assert.Len(t, *pwd, 5)
				err := stringMatchesCharacters(*pwd, characterSet)
				if err != nil {
					t.Errorf("Generate() error = %v", err)
					return
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := password.New(tt.fields.Config)
			got, err := g.GenerateWithLength(tt.args.length)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GenerateWithLength() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.test(got, g.CharacterSet)
		})
	}
}

func TestGenerator_GenerateManyWithLength(t *testing.T) {
	type fields struct {
		Config *domain.PasswordSecurity
	}
	type args struct {
		amount int
		length int
	}
	tests := []struct {
		name    string
		args    args
		fields  fields
		test    func([]string, string)
		wantErr error
	}{
		{
			name:   "valid",
			args:   args{amount: 5, length: 5},
			fields: fields{Config: &password.DefaultConfig},
			test: func(pwds []string, characterSet string) {
				assert.Len(t, pwds, 5)

				for _, pwd := range pwds {
					assert.Len(t, pwd, 5)
					err := stringMatchesCharacters(pwd, characterSet)
					if err != nil {
						t.Errorf("Generate() error = %v", err)
						return
					}
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := password.New(tt.fields.Config)

			got, err := g.GenerateManyWithLength(tt.args.amount, tt.args.length)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("GenerateManyWithLength() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.test(got, g.CharacterSet)
		})
	}
}
