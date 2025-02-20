package method

import (
    "fmt"
    "net/http"
    "encoding/json"
    "math/rand"
    "time"
    // "sync"   -- Doesn't appear to require mutex
)

func Pull(w http.ResponseWriter, r *http.Request){
    // fmt.Fprintf(w,
    //     TempMessage("Request noted!", r),
    //     )

    // switch r.URL.Path:
    // case "mens"

    path := r.URL.Path

    switch {
    case path == "/mens":
        fmt.Fprintf(w, Cubby(M))
    case path == "/womens":
        fmt.Fprintf(w, Cubby(W))
    default:
        fmt.Fprintf(w, "Welcome to the shoe desk\n\n")
        fmt.Fprintf(w, "Please visit either /mens or /womens to see what we have in stock")
    }
}

func Put(w http.ResponseWriter, r *http.Request){

    // fmt.Fprintf(w,
    //     TempMessage("No memory now - sorry!", r),
    // )

    var customer Counter

    // JSON decoding and validation
    err := json.NewDecoder(r.Body).Decode(&customer)

    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Size validation
    right_size := InRange(customer.Shoes)

    if right_size == false{
        msg := fmt.Sprintf("I'm sorry, we don't carry %v size %v shoes", customer.Shoes.Sex, customer.Shoes.Size)
        http.Error(w, msg, http.StatusBadRequest)
        return
    }

    // For now, restricted to home '/' endpoint
    if r.URL.Path == "/" {

        count := ShoeReturn[customer.Shoes]

        switch customer.Request {
        case Return:
            if count > ShoeMax {
                http.Error(w, "Those can't be our shoes! We're maxed out here!\n", http.StatusMethodNotAllowed)
            } else {
                ShoeReturn[customer.Shoes] += 1
                fmt.Fprintf(w, "Thanks for returning your %v size %v shoes!\n", customer.Shoes.Size, customer.Shoes.Sex)
            }

        case Borrow:
            if count == 0 {
                msg := fmt.Sprintf("Sorry, we're out of %v size %v. How about a different size?\n", customer.Shoes.Size, customer.Shoes.Sex)
                http.Error(w, msg, http.StatusMethodNotAllowed)
            } else{
                ShoeReturn[customer.Shoes] -= 1
                fmt.Fprintf(w, "Here are your %v size %v shoes - please return them before end of day!\n", customer.Shoes.Size, customer.Shoes.Sex)
            }
        }
    }
}
