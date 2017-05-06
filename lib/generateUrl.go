package lib

import (
    "time"
    "strings"
    "strconv"
)

type overlayUrl struct {
    TimeNow time.Time
    TimeString string
    Url string
}

func NewOverlayUrl() overlayUrl {
    timeNow := time.Now()
    timeString := getTimeString(timeNow)
    url := getUrl(timeString)
    return overlayUrl{timeNow, timeString, url}
}

func (o *overlayUrl) SubFive() {
    o.TimeNow = o.TimeNow.Add(-5 * time.Minute)
    o.TimeString = getTimeString(o.TimeNow)
    o.Url = getUrl(o.TimeString)
}

func getTimeString(timeNow time.Time) string {
    year, month, date := timeNow.Date()
    hour, min := timeNow.Hour(), timeNow.Minute()
    min = makeFive(min)

    yearStr, monthStr, dateStr, hourStr, minStr := 
    strconv.Itoa(year), strconv.Itoa(date),
    strconv.Itoa(int(month)),
    strconv.Itoa(hour), strconv.Itoa(min)

    yearStr, monthStr, dateStr, hourStr, minStr =
    fixZeroes(yearStr), fixZeroes(monthStr),
    fixZeroes(dateStr), fixZeroes(hourStr), fixZeroes(minStr)

    return yearStr + monthStr + dateStr + hourStr + minStr
}

func fixZeroes(digits string) string {
    if len(digits) == 1 {
        digits = "0" + digits
    }
    return digits
}

func makeFive(num int) int {
    return num - (num % 5)
}

func getUrl(_time string) string {
    urlParts := []string {
        "http://www.weather.gov.sg/files/rainarea/50km/v2/dpsri_70km_", 
        _time, "0000dBR.dpsri.png"}
        return strings.Join(urlParts, "")
}

