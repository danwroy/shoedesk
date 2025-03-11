package method

import (
    "fmt"
    // "encoding/json"
	"strings"
	"unicode"
)

/// String methods

func (s Sex) String() string{
    return defSex[s].String()
}

func (h Handoff) String() string{
    return defHand[h].String()
}

func (m meta) String() string{
    return m.name.String()
}

func (n name) String() string{
    return string(n)
}


/// Input handling

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

// Normalize strings - lowercase, no punct
// (source - ChatGPT)
func norm(s string) string{
	var n strings.Builder
	for _, r := range s {
		if !unicode.IsPunct(r) {
			n.WriteRune(r)
		}
	}
	return strings.ToLower(n.String())
}

// Maps string representation to map object
func reverselookup[K comparable, V fmt.Stringer](m map[K]V) (map[string]K, error){
    // Func argument is a generic map
    // V has to be constrained against fmt.Stringer in order to call .String() on values
    // Also had to add a .String() to var name for this to work

    var r map[string]K

    for k, v := range m {

        name := v.String()
        _, ok := r[norm(name)]      // normalized for lookup

        if !ok {                    // only store if value not in map
            r[norm(name)] = k
        }else{
            return r, fmt.Errorf("Duplicate error: %s occurs more than once", name)
        }
    }
    return r, nil
}
