package main

import (
    "fmt"
    "net/http"
    "nimbus2/lib"
)

func main() {
    overlayUrl := lib.NewOverlayUrl()
    var statusCode int
    response, _ := http.Get(overlayUrl.Url)
    statusCode = response.StatusCode
    fmt.Println(overlayUrl.Url, statusCode)
    for statusCode == 404 {
        overlayUrl.SubFive()
        response, _ = http.Get(overlayUrl.Url)
        statusCode = response.StatusCode
        fmt.Println(overlayUrl.Url, statusCode)
    }
}
