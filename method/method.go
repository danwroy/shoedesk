package method

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func Pull(w http.ResponseWriter, r *http.Request){
    // fmt.Fprintf(w,
    //     TempMessage("Request noted!", r),
    //     )

    // switch r.URL.Path:
    // case "mens"

    // fmt.Fprintf(w, "%v\n\n", r.URL.Path)

    path := r.URL.Path

    switch {
    case path == "/mens":
        fmt.Fprintf(w, Cubby(M))
    case path == "/womens":
        fmt.Fprintf(w, Cubby(W))
    default:
        fmt.Fprintf(w, "Welcome to the shoe counter\n\n")
        fmt.Fprintf(w, "Please visit either /mens or /womens to see what we have in stock")
    }
}

func Put(w http.ResponseWriter, r *http.Request){

    // fmt.Fprintf(w,
    //     TempMessage("No memory now - sorry!", r),
    // )

    var customer Counter
    err := json.NewDecoder(r.Body).Decode(&customer)

    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
    }

    count := ShoeReturn[customer.Shoes]

    switch customer.Request {
    case Return:
        if count > 25 {
            http.Error(w, "Those can't be our shoes! We're maxed out here!", http.StatusMethodNotAllowed)
        } else {
            ShoeReturn[customer.Shoes] += 1
            fmt.Fprintf(w, "Thanks for returning your size %v %v shoes!", customer.Shoes.Size, customer.Shoes.Sex)
        }

    case Borrow:
        if count == 0 {
            fmt.Fprintf(w, "Sorry, we're out of size %v %v. How about a different size?", customer.Shoes.Size, customer.Shoes.Sex)
        } else{
            ShoeReturn[customer.Shoes] -= 1
            fmt.Fprintf(w, "Here are your size %v %v shoes - please return them before end of day!", customer.Shoes.Size, customer.Shoes.Sex)
        }
    }
}
