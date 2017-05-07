package main

import (
    "fmt"
    "nimbus2/lib"
)

func main() {
    download := lib.NewDownload()
    fmt.Println(download.Url, download.StatusCode)
    download.SubFiveMins()
    fmt.Println(download.Url, download.StatusCode)
}
