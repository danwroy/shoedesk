package method


//// Doing an enum approximate using golang patterns
// A. Define as types (int for iota definition)
type Sex int
type Exchange int

// B. Define as constants with iota
// (from 1 so that 0 can indicate error)
const (
    M Sex = iota + 1
    W
)
const (
    borrow Exchange = iota + 1
    return
)


//// Define objects
type Shoes struct {
    Size        int         `json:"size"`
    Sex         Sex         `json:"sex"`
}

type Customer struct {
    Exchange    Exchange    `json:"exchange"`
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
type x bool


var (

    ShoeReturn      map[Shoes]int = make(map[Shoes]int)                         // In-memory shoe return
    paramsSex       map[Sex]def[sizes] = make(map[Sex]def[sizes])               // string repr + shoe size range
    paramsExchange  map[Exchange]def[x] = make(map[Exchange]def[x])             // string repr only

)
