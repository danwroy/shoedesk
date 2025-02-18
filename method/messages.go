package method

import (
    "fmt"
    "net/http"
)

func TempMessage(m string, r *http.Request) string {
    return fmt.Sprintf("%s\nMethod: %v\nURL: %v%v\n", m, r.Method, r.Host, r.URL.Path)
}
