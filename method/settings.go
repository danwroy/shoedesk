package method

// Basic system-wide settings below

var (

    CubbyWidth = 5                  // How many columns across will the cubby be
    ShoeMax = 25                    // Maximum number of shoes per size

)

// Define parameters for key constants
var (

    // String representation for Sex constants
    M_name = "men's"
    W_name = "women's"

    // Shoe size range
    M_sizes = sizes{4, 16}
    W_sizes = sizes{5, 13}

    // String representation for Exchange constants
    borrow_name = "borrow"
    return_name = "return"

)


// Above settings initialized at runtime
func init() {

    // Set parameters per shoe type (by Sex)
    // {string repr, max size, min size}
    paramsSex[M] = def[sizes]{name: M_name, set:M_sizes}
    paramsSex[W] = def[sizes]{name: W_name, set:W_sizes}

    // Set parameters per Exchange type (name only)
    paramsExchange[borrow] = def[x]{name: borrow_name}       // The "set" field should return the empty string
    paramsExchange[return] = def[x]{name: return_name}

}
