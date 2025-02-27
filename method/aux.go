package method

import (
    // "fmt"
    "time"
    "math/rand"
    "encoding/json"
    // "strconv"
    "log"
)

// To validate shoe size by comparing against rage per sex
func InRange(shoe Shoes) bool{
    switch {
    case shoe.Size < Limits[shoe.Sex].Min:
        return false
    case shoe.Size > Limits[shoe.Sex].Max:
        return false
    default:
        return true
    }
}

// Initialize cubby with random numbers
func CubbyStart(max int){
    men := Limits[M].Max
    women := Limits[W].Max
    rand.Seed(time.Now().UnixNano())

    for size := men; size >= Limits[M].Min; size--{
        ShoeReturn[Shoes{size, M}] = rand.Intn(max)
    }
    for size := women; size >= Limits[W].Min; size--{
        ShoeReturn[Shoes{size, W}] = rand.Intn(max)
    }
}

// func (c Counter) MarshalJSON() ([]byte, error) {
//
//     inlined := type struct {Request, Size, Sex}
//
//     inlined {c.Request, c.Shoes.Size, c.Shoes.Sex}
//
// 	return json.Marshal(s)
// }

func (c *Counter) UnmarshalJSON(b []byte) error {

    var s map[string]any

    err := json.Unmarshal(b, &s)
    if err != nil {
        log.Fatal(err)
    }

    request, size, sex := s["request"], s["size"], s["sex"]

    c.Request, c.Shoes.Size, c.Shoes.Sex = request.(Request), size.(int), sex.(Sex)
    // The <var>.(<type>) is type assertion; required by engine

    return nil
}
