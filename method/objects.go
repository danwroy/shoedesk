package method

// import "sync"

type Shoes struct {
    Size    int             `json:"size"`
    Sex     Sex             `json:"sex"`
}

type Counter struct {
    Request     Request     `json:"request"`
    Shoes       Shoes       `json:"shoes"`
}

type Limit struct {
    Min int                 `json:"min"`
    Max int                 `json:"max"`
}

type Sex string
type Request string

const (
    M Sex = "mens"
    W Sex = "womens"
    Borrow Request = "borrow"
    Return Request = "return"
)

var (
	ShoeReturn  map[Shoes]uint8 = make(map[Shoes]uint8) // Create dict representing shoe return

    Limits      map[Sex]Limit = make(map[Sex]Limit)
	// Mu          sync.Mutex
)
