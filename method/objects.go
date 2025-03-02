package method


//// Doing an enum approximate using golang patterns
// A. Define as types (int for iota definition)
type Sex int
type Exchange int

// B. Define as constants with iota
// (from 1 so that 0 can indicate error)
const (
    M Sex = iota + 1
    W Sex
)
const (
    Borrow Exchange = iota + 1
    Return Exchange
)


//// Define objects

type Shoes struct {
    Size    int             `json:"size"`
    Sex     string          `json:"sex"`
    // note - defined differently than Sex type
}

type Customer struct {
    Exchange    Exchange    `json:"exchange"`
    Shoes
}

// Special metadata object for shoe type (sex)
type Def struct {
    Name    string          // string representation of sex
    Min     int             // max size of shoe
    Max     int             // min size
}


// Define in-memory storage items
var (

    ShoeReturn      map[Shoes]int = make(map[Shoes]int)             // In-memory shoe return
    Params          map[Sex]Def = make(map[Sex]Def)                 // Defined at settings.go

)

