package method

import (
    // "fmt"
    "time"
    "math/rand"
    "encoding/json"
    // "strconv"
    "log"
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

// Initialize cubby with random numbers
func CubbyStart(max int){

    rand.Seed(time.Now().UnixNano())

    for sex, p in Params{

        size := p.Max

        for size := max; size >= p.Min; size-- {
            ShoeReturn[Shoes{size, sex}] = rand.Intn(max)
        }
    }
}
