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
func (c *Constants) String() string{
    return Params[c].Name
}

// Unmarshaler to convert 
func (c Customer) UnmarshalJSON(b []byte) {

    input := struct{
        exchange    string
        size        int
        sex         string
    }

    if err := json.Unmarshal(b, &input); err != nil{
        return err
    }

    return Customer{input.exchange, Shoes{input.size, input.sex}}
}
