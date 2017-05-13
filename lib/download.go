package lib

import (
    "io"
    "os"
    "time"
    "strings"
    "net/http"
)

type download struct {
    Time time.Time
    TimeString string
    Url string
    FileName string
    Response *http.Response
    StatusCode int
}

func NewDownload() download {
    _download := download{}
    _download.Update(time.Now())
    return _download
}

func (d *download) Update(_time time.Time) {
    d.Time = _time
    d.TimeString = getTimeString(_time)
    d.Url = getUrl(d.TimeString)
    d.FileName = d.TimeString + ".png"
    d.Response, _ = http.Get(d.Url)
    d.StatusCode = d.Response.StatusCode
}

func (d *download) SubFiveMins() {
    d.Update(d.Time.Add(-5 * time.Minute))
}

func (d *download) Save() {
    defer d.Response.Body.Close()
    CheckAndCreate("./records")

    file, err := os.Create("./records/" + d.FileName)
    if err != nil { panic(err) }
    defer file.Close()

    _, err = io.Copy(file, d.Response.Body)
    if err != nil { panic(err) }
}

func getTimeString(_time time.Time) string {
    _time = _time.Round(5 * time.Minute)
    return _time.Format(
        "2006" + "01" + "02" + "15" + "04")
}

func getUrl (_timeString string) string {
    urlSlice := []string {
        "http://www.weather.gov.sg/files/rainarea/50km/v2/dpsri_70km_", 
        _timeString, "0000dBR.dpsri.png"}
    return strings.Join(urlSlice, "")
}
