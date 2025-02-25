package method

type Shoes struct {
    Size    int             `json:"size"`
    Sex     Sex             `json:"sex"`
}

type Counter struct {
    Request Request          `json:"request"`
    Shoes   Shoes            `json:inline`
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
	ShoeReturn  map[Shoes]int = make(map[Shoes]int)     // Create dict representing shoe return
    Limits      map[Sex]Limit = make(map[Sex]Limit)     // Defined @ settings.go
)
