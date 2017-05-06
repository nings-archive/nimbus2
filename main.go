package main

import (
    "fmt"
    "nimbus2/lib"
)

func main() {
    overlayUrl := lib.NewOverlayUrl()
    fmt.Println("Now:", overlayUrl.Url)
    overlayUrl.SubFive()
    fmt.Println("Sub:", overlayUrl.Url)
}
