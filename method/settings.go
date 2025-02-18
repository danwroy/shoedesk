package method

var RowLength int = 5               // how many columns across will the cubby be
var ShoeMax int = 25

func init() {                       // define largest/smallest sides
    Limits[M] = Limit{4, 16}        // men's
    Limits[W] = Limit{5, 13}        // women's
}                                   // "Limits" defined @ objects.go
