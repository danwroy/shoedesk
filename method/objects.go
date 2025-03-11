package method


//// Doing an enum approximate using golang patterns

// A. Define as types (int for iota definition)
type Sex int
type Handoff int

// B. Define as constants with iota
const (
    M Sex = iota + 1                // to free up 0 for error
    W
)
const (
    Borrow Handoff = iota + 1
    Return
)

//// Define objects

type Shoes struct {
    Size        int         `json:"size"`
    Sex         Sex         `json:"sex"`
}

type Customer struct {
    Handoff     Handoff     `json:"handoff"`
    Shoes                   // Has to be defined w/ type def for JSON inline requests
}

type name string

type sizes struct {
    min int
    max int
}

type meta struct {          // Will hold info about shoes per sex
    name
    sizes
}

// Mappings

var (

    ShoeReturn      map[Shoes]int       = make(map[Shoes]int)       // In-memory shoe return
    defSex          map[Sex]meta        = make(map[Sex]meta)        // string repr + shoe size range
    defHand         map[Handoff]name    = make(map[Handoff]name)    // string repr only
    lookupSex       map[string]Sex      = make(map[string]Sex)      // Necessary string lookup map
    lookupHand      map[string]Handoff  = make(map[string]Handoff)  // ''

)
