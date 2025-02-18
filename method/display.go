package method

import (
    "fmt"
    // "strings"
)

func Cubby(sex Sex) string {
    min := Limits[sex].Min          // see: settings.go
    max := Limits[sex].Max
    // min := 4
    // max := 16
    col := 0                        // iterated to find cubby width
    cubby := ""                     // cubby is a text object

    for size := min; size <= max; size++ {

        stock := ShoeReturn[Shoes{size, sex}]

        cubby += fmt.Sprintf("| %2d (%-2d pairs) |", size, stock)
        
        // row break set if RowLength met
        col += 1
        if col % RowLength == 0 {
            cubby += "\n"
        }
    }
    return cubby
}
