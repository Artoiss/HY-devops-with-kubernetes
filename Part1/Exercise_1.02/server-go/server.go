package main

import (
    "fmt"
    "net/http"
    "log"
)

func main() {
    fmt.Printf("Server started in port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
