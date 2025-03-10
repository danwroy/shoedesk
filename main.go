package main

import (
    "fmt"
    "net/http"
    m "shoedesk/method"
)

func main() {

    // initialize cubby
    m.FillReturn(m.ShoeMax)

    // using system standard muxer
    route := http.NewServeMux()


    route.HandleFunc("GET /", get)
    route.HandleFunc("PATCH /", patch)
    route.HandleFunc("GET /{sex}", get)
    route.HandleFunc("PATCH /{sex}", patch)
    route.HandleFunc("GET /{sex}/{size}", get)
    route.HandleFunc("PATCH /{sex}/{size}", patch)

    server := http.Server{
        Addr: ":8080",
        Handler: route
    }

    // Activate server
    fmt.Println("Server listening on port ", server.Addr)
    server.ListenAndServe()
}

// define how http requests are handled



// func get(w http.ResponseWriter, r *http.Request){
//     switch r.Method {
//
//     case "GET":
//         m.Pull(w, r)
//
//     case "PATCH":
//         m.Update(w, r)
//     }
// }
