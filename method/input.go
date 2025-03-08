package method

import (
    // "fmt"
    "encoding/json"
)

/// Input processing

// To validate shoe size by comparing against rage per sex
func InRange(shoe Shoes) bool{
    sex, size := shoe.Sex, shoe.Size 
    min, max := Params[sex].min, Params[sex].max
    if min <= size && size <= max {
        return true
    } else {
        return false
    }
}

// Return string for sex (set at settings.go)
func (c *Sex) String() string{
    return paramsSex[c].name
)

func (c *Exchange) String() string{
    return paramsExchange[c].name
}
