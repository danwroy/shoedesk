package method

// Basic system-wide settings below

var (

    CubbyWidth = 5                  // How many columns across will the cubby be
    ShoeMax = 25                    // Maximum number of shoes per size

)

// Define enum constants at objects.go
var (

    // String representation for Sex constants
    M_name = "men's"
    W_name = "women's"

    // Shoe size range per Sex
    M_sizes = sizes{4, 16}
    W_sizes = sizes{5, 13}

    // String representation for Handoff constants
    borrow_name = "borrow"
    return_name = "return"

)


// Above settings initialized at runtime
// Be sure to add any necessary initialization process here
func init() {

    // Set parameters per shoe type (by Sex)
    // {string repr, max size, min size}
    defSex[M] = def[sizes]{name: norm(M_name), set:M_sizes}
    defSex[W] = def[sizes]{name: norm(W_name), set:W_sizes}

    // Set parameters per Handoff type (name only)
    defHandoff[borrow] = def[x]{name: norm(borrow_name)}         // The "set" field here should return the empty string
    defHandoff[return] = def[x]{name: norm(return_name)}

}
