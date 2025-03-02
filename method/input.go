package method

import (
    // "fmt"
    "time"
    "math/rand"
    "encoding/json"
    // "strconv"
    // "log"
)

/// Input processing

// To validate shoe size by comparing against rage per sex
func InRange(shoe Shoes) bool{
    switch {
    case shoe.Size < Params[shoe.Sex].Min:
        return false
    case shoe.Size > Limits[shoe.Sex].Max:
        return false
    default:
        return true
    }
}

/// Convert string values to proper value
func (e Exchange) String() string{
    return Params[e].Name
}

func (s Sex) String() string{
    return Params[s].Name
}
