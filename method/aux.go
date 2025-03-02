package method

import (
    // "fmt"
    "time"
    "math/rand"
    "encoding/json"
    // "strconv"
    "log"
)

/// Array functions

// return index number for string in array
// func (a Parameters) Index(s string) int {
//     for i, v := range a {
//         if v == s {
//             return i
//         }
//     return nil
//     }
// }

// Print all strings in an array
// func ShowAll(a Parameters) string {
//     for i in a {
//         fmt.Println(i)
//     }
//     return
// }


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

// func (c Counter) MarshalJSON() ([]byte, error) {
//
//     inlined := type struct {Request, Size, Sex}
//
//     inlined {c.Request, c.Shoes.Size, c.Shoes.Sex}
//
// 	return json.Marshal(s)
// }

// func (c *Counter) UnmarshalJSON(b []byte) error {
//
//     var s map[string]any
//
//     err := json.Unmarshal(b, &s)
//     if err != nil {
//         log.Fatal(err)
//     }
//
//     request, size, sex := s["request"], s["size"], s["sex"]
//
//     c.Request, c.Shoes.Size, c.Shoes.Sex = request.(Request), size.(int), sex.(Sex)
//     // The <var>.(<type>) is type assertion; required by engine
//
//     return nil
// }
