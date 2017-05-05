package main

import (
    "fmt"
    "nimbus2/lib"
)

func main() {
    overlay_url := lib.NewOverlayUrl()
    fmt.Println(overlay_url.Url)
}
