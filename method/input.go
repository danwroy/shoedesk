package method

import (
    // "fmt"
    "encoding/json"
)

/// Input processing

// To validate shoe size by comparing against rage per sex
func InRange(shoe Shoes) bool{
    sex, size := shoe.Sex, shoe.Size 
    min, max := Params[sex].Min, Params[sex].Max
    if min <= size && size <= max {
        return true
    } else {
        return false
    }
}

// Return string for sex (set at settings.go)
func (s Sex) String() string{
    return Params[s].Name
}

// Return string for Exchange type
func (e Exchange) String() string{
    switch e {
    case Borrow:
        return "borrow"
    case Return:
        return "return"
    default:
        return ""
    }
}

// const (
//
// ArgsM Args = Args{
//         "men's", "mens", "man's", "mans", "man", "male", "boy's", "boys", "guy's", "guys", "m"
//     }
// ArgsW Args = Args{
//         "men's", "womens", "woman's", "womans", "woman", "female", "girl", "girl's", "girls", "ladies", "lady's", "w", "f", "l"
//     }
// ArgsBorrow Args = Args{
//         "borrow", "rent", "take", "get", "lift"
//     }
// ArgsReturn Args = Args{
//         "return", "leave", "give", "yield", "bring"
//     }
//
// )

// Unmarshaler to convert 
func (c *Customer) UnmarshalJSON(b []byte) error{

    input := struct{
        exchange: string
        size: int
        sex: string
    }

    if err := json.Unmarshal(b, &input); err != nil{
        return err
    }

}
