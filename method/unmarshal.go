package method

import (
    "encoding/json"
    "fmt"
)

func (sx *Sex) UnmarshalText(text byte[]) {
    switch string(text){
    case M.String():
        *sx = M
    case W.String():
        *sx = W
    default:
        return fmt.Errorf("'%v' not valid: '%s' or '%s' please", text, M, W)
    }
}

func (e *Handoff) UnmarshalText(text byte[]){
    switch string(text){
    case borrow.String():
        *e = borrow
    case return.String():
        *e = return
    default:
        return fmt.Errorf("'%v' not valid: '%s' or '%s' please", text, borrow, return)
    }
}
