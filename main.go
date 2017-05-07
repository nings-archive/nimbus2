package main

import (
    "fmt"
    "nimbus2/lib"
)

func main() {
    overlayUrl := lib.NewOverlayUrl()
    download := lib.NewDownload(&overlayUrl)
    fmt.Println(download.OverlayUrl.TimeString, download.StatusCode)
    download.SubFive()
    fmt.Println(download.OverlayUrl.TimeString, download.StatusCode)
}
