package method

import (
    // "fmt"
    "encoding/json"
	"strings"
	"unicode"
)

/// Input processing

// To validate shoe size by comparing against rage per sex
func InRange(shoe Shoes) bool{
    sex, size := shoe.Sex, shoe.Size 
    min, max := defSex[sex].min, defSex[sex].max
    if min <= size && size <= max {
        return true
    } else {
        return false
    }
}

// Return string for sex (set at settings.go)
func (c *Sex) String() string{
    return defSex[c].name
)

func (c *Handoff) String() string{
    return defHandoff[c].name
}

// Normalize strings
func norm(s string) string {
	var n strings.Builder
	for _, r := range s {
		if !unicode.IsPunct(r) {
			n.WriteRune(r)
		}
	}
	return strings.ToLower(n.String())
}
