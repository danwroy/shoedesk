package method

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io"
    "log"
    "strconv"
)

// GET methods
func Pull(w http.ResponseWriter, r *http.Request){

    // "PathValue" returns {path param} value from muxer
    sex  := lookupSex[ norm(r.PathValue("sex")) ]                   // Both should return 0 if no PathValue
    size, _ := strconv.Atoi(r.PathValue("size"))


    // Main conditionals

    // print statement
    fmt.Println(r.URL.Path)

    if r.URL.Path == "/"{
        fmt.Fprintf(w, "Welcome to the shoe desk\n\n")
        fmt.Fprintf(w, "See what we have available by checking by /<sex> " +
            "or by /<sex>/<size>")
    }

    if sex != 0{

        if defSex[sex].min <= size && size <= defSex[sex].max {
            fmt.Fprintf(w, Cubby(sex, sizes{size, size}))           // Show cubby hole for one set of shoes

        }else{
            fmt.Fprintf(w, Cubby(sex, defSex[sex].sizes) )          // Show cubby for all of one sex
        }
    }
}


// PATCH methods - borrow and return
func Update(w http.ResponseWriter, r *http.Request){

    sex  := lookupSex[ norm(r.PathValue("sex")) ]
    size, _ := strconv.Atoi(r.PathValue("size"))

    customer := Customer{}

    // Use values from path parameters first, if available (will be overwritten by direct request)
    if size != 0{
        customer.Size = size
    }
    if sex != 0{
        customer.Sex = sex
    }

    // print statement
    // fmt.Println(customer)

    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Fatalf("Issue decoding body:", err)
    }

    // JSON decoding and validation
    err = json.Unmarshal(body, &customer)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        fmt.Println(err)
        return
    }

    // (print statement)
    // fmt.Println(customer)

    // Size validation 
    if InRange(customer.Shoes) == false{
        msg := fmt.Sprintf("I'm sorry, we don't keep %s size %d shoes", 
            customer.Shoes.Sex.String(), 
            customer.Shoes.Size)
        http.Error(w, msg, http.StatusBadRequest)
        return
    }


    // Main conditional

    stock := ShoeReturn[customer.Shoes]

    switch customer.Handoff {
    case Return:
        if stock > ShoeMax {
            http.Error(w, "Those can't be our shoes! We're maxed out here!\n", http.StatusMethodNotAllowed)
            // Newline added to eliminate trailing % sign
        } else {
            ShoeReturn[customer.Shoes] += 1
            fmt.Fprintf(w, "Thanks for returning your size %d %s shoes!\n", customer.Shoes.Size, customer.Shoes.Sex.String())
        }

    case Borrow:
        if stock == 0 {
            msg := fmt.Sprintf("Sorry, we're out of size %d %s. How about a different size?\n", customer.Shoes.Size, customer.Shoes.Sex.String())
            http.Error(w, msg, http.StatusMethodNotAllowed)
        } else{
            ShoeReturn[customer.Shoes] -= 1
            fmt.Fprintf(w, "Here are your size %d %s shoes - please return them before end of day!\n", customer.Shoes.Size, customer.Shoes.Sex.String())
        }
    }
}

