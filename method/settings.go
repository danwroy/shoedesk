package method

var (

    CubbyWidth = 5                    // How many columns across will the cubby be
    ShoeMax = 25                    // Maximum number of shoes per size
    Range_M = SizeRange{4, 16}
    Range_W = SizeRange{5, 13}

)


// Settings initialized at runtime
func init() {

    // Set parameters per shoe type (by Sex)
    // {string repr, max size, min size}
    Params[M] = Def[SizeRange]{Name:"men's", Set:Range_M}
    Params[W] = Def[SizeRange]{Name:"women's", Set:Range_W}

    // Set parameters per Exchange type (name only)
    Params[Borrow] = Def[string]{Name: "borrow"}
    Params[Return] = Def[string]{Name: "return"}

    // Map name strings to constants
    for k, v := range Params{
        GetConst[v.Name] = k
    }

}
