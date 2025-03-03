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
    Borrow Exchange = iota + 1
    Return
)

// Allows combining operations on both
type Constants interface {
    Sex | Exchange
}


//// Define objects

type Shoes struct {
    Size        int         `json:"size"`
    Sex         Sex         `json:"sex"`
}

type Customer struct {
    Exchange    Exchange    `json:"exchange"`
    Shoes
}


// Special metadata object for shoe type (sex)

type Def[T any] struct {
    Name    string          // string representation of sex
    Set     T               // allows optional settings to be mapped
}

// [Set] option
type SizeRange struct {
    Min int
    Max int
    }


// All terms for M/W
type Args = [...]string

var (

    ShoeReturn      map[Shoes]int = make(map[Shoes]int)                     // In-memory shoe return
    Params          map[Constants]Def = make(map[Constants]Def)                   // Defined at settings.go
    StringRep       map[string]Constants = make(map[string]Constants)       // Map strings to constants

)

