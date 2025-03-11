package method

import (
    // "encoding/json"
    "fmt"
)

func (sex *Sex) UnmarshalText(text []byte){
    switch string(text){
    case M.String():
        *sex = M
    case W.String():
        *sex = W
    default:
        fmt.Errorf("'%v' not valid: '%s' or '%s' please", text, M, W)
    }
}

func (h *Handoff) UnmarshalText(text []byte){
    switch string(text){
    case Borrow.String():
        *h = Borrow
    case Return.String():
        *h = Return
    default:
        fmt.Errorf("'%v' not valid: '%s' or '%s' please", text, Borrow, Return)
    }
}
