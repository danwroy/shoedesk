package method

import (
    "fmt"
    "time"
    "math/rand"
)

// To display a representation of a shoe return cubby
func Cubby(sex Sex) string {
    min := Params[sex].Min          // see: settings.go
    max := Params[sex].Max
    col := 0                        //  iterated to find cubby width
    cubby := "||"                   // cubby is a text object; lead with outer formatting

    for size := min; size <= max; size++ {

        stock := ShoeReturn[Shoes{size, sex}]

        cubby += fmt.Sprintf(" %2d (%-2d pairs) |", size, stock)
        
        // row break set if RowLength met
        col += 1
        if col % CubbyWidth == 0 {
            cubby += "|\n||"
        }
    }
    return cubby
}

// Initialize cubby with random numbers
func FillReturn(max int){

    rand.Seed(time.Now().UnixNano())

    for c, p := range Params{

        if c == Sex {

            sex := c
            for size := p.Min; size < p.Max; size++ {
                ShoeReturn[Shoes{size, sex}] = rand.Intn(max)
            }
        }
    }
}
