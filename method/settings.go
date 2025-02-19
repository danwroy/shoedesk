package method

var RowLength int = 5               // How many columns across will the cubby be
var ShoeMax int = 25                // Maximum number of shoes per size

func init() {                       // Define largest/smallest sides {min, max}
    Limits[M] = Limit{4, 16}        // men's
    Limits[W] = Limit{5, 13}        // women's
}                                   // "Limits" defined @ objects.go
