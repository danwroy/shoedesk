package method

type Shoes struct {
    Size    int             `json:"size"`
    Sex     int             `json:"sex"`
}

type Counter struct {
    Request boolean         `json:"request"`
    Shoes   Shoes           `json:",inline"`
}

type Limit struct {
    Min int                 `json:"min"`
    Max int                 `json:"max"`
}

type Sex int
type Request int

const (
    M Sex = iota
    W Sex
)

const (
    Borrow Request = iota
    Return Request
)

// return string for Sex value
func (sx Sex) String() string {
    return [...]{"men's", "women's"}[sx]
}

// return string for Request value
func (rq Request) String() string (
    return [...]string{"borrow", "return"}[rq]
)

var (
    ShoeReturn  map[Shoes]int = make(map[Shoes]int)                                             // Create dict representing shoe return
    Limits      map[string]Limit = make(map[string]Limit)                                       // Defined @ settings.go
    // SexStr      map[boolean]string = map[boolean]string{true: "men's", false: "women's"})   // Map boolean to sex
    // RequestStr  map[boolean]string = map[boolean]string{true: "borrow", false:"return"})    // Map boolean to request type
)

