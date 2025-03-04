package method

import (
    "encoding/json"
    "fmt"
)

func (sx *Sex) UnmarshalText(text byte[]) {
    switch string(text){
    case "men's":
        *sx = M
    case "women's":
        *sx = W
    default:
        return fmt.Errorf("'%v' not valid: '%s' or '%s' please", text, M, W)
    }
}

func (e *Exchange) UnmarshalText(text byte[]){
    switch string(text){
    case "borrow":
        *e = Borrow
    case "return":
        *e = Return
    default:
        return fmt.Errorf("'%v' not valid: '%s' or '%s' please", text, Borrow, Return)
    }
}
