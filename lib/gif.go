package lib

import (
    "os"
    "image"
    "image/png"
    "image/gif"
)

func MakeGif(files []string) *gif.GIF {
    outGif := &gif.GIF{}
    for _, name := range files {
        frame := pngToGif(name)
        outGif.Image = append(outGif.Image, frame.(*image.Paletted))
        outGif.Delay = append(outGif.Delay, 0)
    }
    os.Remove("temp.gif")
    return outGif
}

func pngToGif(name string) image.Image {
    file, err := os.Open("./records/" + name)
    if err != nil { panic(err) }
    png, err := png.Decode(file)
    if err != nil { panic(err) }
    file.Close()

    if err != nil { panic(err) }
    file, err = os.OpenFile("temp.gif", os.O_WRONLY|os.O_CREATE, 0444)
    err = gif.Encode(file, png, &gif.Options{256, nil, nil})
    if err != nil { panic(err) }
    file.Close()

    file, err = os.Open("temp.gif")
    defer file.Close()
    if err != nil { panic(err) }
    gif, err := gif.Decode(file)
    if err != nil { panic(err) }

    return gif
}

func SaveGif(outGif *gif.GIF) {
    file, err := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0777)
    if err != nil { panic(err) }
    defer file.Close()
    err = gif.EncodeAll(file, outGif)
    if err != nil { panic(err) }
}
