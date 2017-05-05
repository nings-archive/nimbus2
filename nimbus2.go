package main

import (
    "fmt"
    "nimbus2/lib"
)

func main() {
    overlay_url := lib.NewOverlayUrl()
    fmt.Println("Now:", overlay_url.Url)
    overlay_url.SubFive()
    fmt.Println("Sub:", overlay_url.Url)
}
