// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package code

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseISSUE_BYFromString Parse ISSUE_BY from string
func ParseISSUE_BYFromString(str string) (ISSUE_BY, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := ISSUE_BY_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown ISSUE_BY: %s", str)
	}

	return ISSUE_BY(v), nil
}

// Equal type compare
func (t ISSUE_BY) Equal(target ISSUE_BY) bool {
	return t == target
}

// IsIn todo
func (t ISSUE_BY) IsIn(targets ...ISSUE_BY) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t ISSUE_BY) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *ISSUE_BY) UnmarshalJSON(b []byte) error {
	ins, err := ParseISSUE_BYFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
