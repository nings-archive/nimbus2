package main

import (
    "fmt"
    "nimbus2/lib"
)

func main() {
    download := lib.NewDownload()
    fmt.Println(download.Url, download.StatusCode)
    for download.StatusCode == 404 {
        download.SubFiveMins()
        fmt.Println(download.Url, download.StatusCode) 
    }

    lib.CheckAndCreate("./testdir")
}
