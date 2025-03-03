package main

import (
    "net/http"
    "shoedesk/method"
)

func main() {

    // initialize cubby
    method.FillReturn(method.ShoeMax)

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

// define how http requests are handled
func handler(w http.ResponseWriter, r *http.Request){
    switch r.Method {

    case "GET":
        method.Pull(w, r)

    case "PATCH":
        method.Update(w, r)
    }
}
