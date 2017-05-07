package lib

import (
    "net/http"
)

type download struct {
    OverlayUrl overlayUrl
    Response http.Response
    StatusCode int
}

func NewDownload(o *overlayUrl) download {
    response, _ := http.Get(o.Url)
    return download{*o, *response, response.StatusCode}
}

func (d *download) SubFive() {
    d.OverlayUrl.SubFive()
    response, _ := http.Get(d.OverlayUrl.Url)
    d.Response = *response
    d.StatusCode = d.Response.StatusCode
}
