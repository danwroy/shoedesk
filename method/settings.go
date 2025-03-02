package method

var (

    RowWidth int = 5                    // How many columns across will the cubby be
    ShoeMax int = 25                    // Maximum number of shoes per size

)


// Settings initialized at runtime
func init() {

    // Set parameters per shoe type (sex)
    // {string repr, max size, min size}
    Params[M] = Def{"men's", 4, 16}
    Params[W] = Def{"women's", 5, 13}

}
