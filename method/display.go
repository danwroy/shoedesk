package method

import (
    "fmt"
    "time"
    "math/rand"
)

// To display a representation of a shoe return cubby
func Cubby(sex Sex, sizes sizes) string {
    min := sizes.min
    max := sizes.max
    col := 0                // Iterated to find cubby width
    cubby := "||"           // Initial string to build cubby

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
func FillReturn(cap int){

    rand.Seed(time.Now().UnixNano())

    for sex, r := range defSex{

        for size := r.min; size < r.max; size++ {
            ShoeReturn[Shoes{size, sex}] = rand.Intn(cap)
        }
    }
}
