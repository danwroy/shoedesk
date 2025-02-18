package method

// import "sync"

type Shoes struct {
    Size    int             `json:"size"`
    Sex                     `json:"sex"`
}

type Counter struct {
    Request                 `json:"request"`
    Shoes                   `json:"shoes,inline"`
}

type Limit struct {
    Min int                 `json:"min"`
    Max int                 `json:"max"`
}

type Sex string             //`json:"sex"`
type Request string         //`json:"request"`

const (
    M Sex = "mens"
    W Sex = "womens"
    Borrow Request = "borrow"
    Return Request = "return"
)

var (
	ShoeReturn  map[Shoes]uint8 = make(map[Shoes]uint8)     // Create dict representing shoe return

    Limits      map[Sex]Limit = make(map[Sex]Limit)         // Defined @ settings.go
	// Mu          sync.Mutex
)
