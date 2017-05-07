package lib

import (
    "net/http"
)

type download struct {
    OverlayUrl *overlayUrl
    Response *http.Response
}

func NewDownload(o *overlayUrl) download {
    response, _ := http.Get(o.Url)
    return download{o, response}
}

func (d *download) SubFive() {
    d.OverlayUrl.SubFive()
    response, _ := http.Get(d.OverlayUrl.Url)
    d.Response = response
}
