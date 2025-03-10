package method


//// Doing an enum approximate using golang patterns
// A. Define as types (int for iota definition)
type Sex int
type Handoff int
// type name string

// B. Define as constants with iota
// (from 1 so that 0 can indicate error)
const (
    M Sex = iota + 1
    W
)
const (
    borrow Handoff = iota + 1
    return
)


//// Define objects
type Shoes struct {
    Size        int         `json:"size"`
    Sex         Sex         `json:"sex"`
}

type Customer struct {
    Handoff    Handoff      `json:"handoff"`
    Shoes
}


// Special metadata object
// Any kind of [set] option can be defined
// Define metadata at settings.go file
type def[T any] struct {
    name    string          // string representation
    set     *T              // any chosen additional settings
}

// To define shoe range
type sizes struct {
    min int
    max int
}

// A special type/value to represent nothing
type x *int


var (

    ShoeReturn      map[Shoes]int = make(map[Shoes]int)             // In-memory shoe return
    defSex          map[Sex]def[sizes] = make(map[Sex]def[sizes])   // string repr + shoe size range
    defHandoff      map[Handoff]def[x] = make(map[Handoff]def[x])   // string repr only

)
