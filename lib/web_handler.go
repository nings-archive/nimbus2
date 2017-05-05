package lib

import (
"time"
"strings"
"strconv"
)

type overlay_url struct {
    Time_now string
    Url string
}

func NewOverlayUrl() overlay_url {
    time_now := get_time()
    url := get_url(time_now)
    return overlay_url{time_now, url}
}

func get_time() string {
    time_now := time.Now()
    year, month, date := time_now.Date()
    hour, min := time_now.Hour(), time_now.Minute()
    min = make_five(min)
    year_str, month_str, date_str, hour_str, min_str := 
    strconv.Itoa(year), strconv.Itoa(date),
    strconv.Itoa(int(month)),
    strconv.Itoa(hour), strconv.Itoa(min)
    year_str, month_str, date_str, hour_str, min_str =
    fix_zeroes(year_str), fix_zeroes(month_str),
    fix_zeroes(date_str), fix_zeroes(hour_str), fix_zeroes(min_str)
    return year_str + month_str + date_str + hour_str + min_str
}

func fix_zeroes(digits string) string {
    if len(digits) == 1 {
        digits = "0" + digits
    }
    return digits
}

func make_five(num int) int {
    return num - (num % 5)
}

func get_url(_time string) string {
    url_parts := []string {
        "http://www.weather.gov.sg/files/rainarea/50km/v2/dpsri_70km_", 
        _time, "0000dBR.dpsri.png"}
        return strings.Join(url_parts, "")
}

