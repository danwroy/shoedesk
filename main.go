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

    // All routed methods
    route.HandleFunc("GET /", m.Pull)
    route.HandleFunc("PATCH /", m.Update)
    route.HandleFunc("GET /{sex}/", m.Pull)
    route.HandleFunc("PATCH /{sex}/", m.Update)
    route.HandleFunc("GET /{sex}/{size}/", m.Pull)
    route.HandleFunc("PATCH /{sex}/{size}/", m.Update)

    server := http.Server{
        Addr: ":8080",
        Handler: route,
    }

    // Activate server
    fmt.Println("Server listening on port ", server.Addr)
    server.ListenAndServe()
}
