package method


// Basic system-wide settings below
var (

    CubbyWidth = 5                  // How many columns across will the cubby be
    ShoeMax = 25                    // Maximum number of shoes per size

)

// Define enum constants at objects.go
var (

    // String representation for Sex constants
    M_name name = "men's"
    W_name name = "women's"

    // Shoe size range per Sex
    M_sizes sizes = sizes{4, 16}
    W_sizes sizes = sizes{5, 13}

    // String representation for Handoff constants
    Borrow_name name = "borrow"
    Return_name name = "return"

)


// Above settings initialized at runtime
// Be sure to add any necessary initialization process here
func init() {

    // Set parameters per shoe type (by Sex)
    // {string repr, max size, min size}
    defSex[M] = meta{name: M_name, sizes: M_sizes}
    defSex[W] = meta{name: W_name, sizes: W_sizes}

    // Set parameters per Handoff type (name only)
    defHand[Borrow] = Borrow_name 
    defHand[Return] = Return_name        

    // Initialize string lookup (reverse dictionaries)
    lookupSex, _   = reverselookup(defSex)
    lookupHand, _  = reverselookup(defHand)

}
