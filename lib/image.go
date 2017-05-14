package lib

import (
    "os"
    "image"
    "image/png"
    "image/draw"
    "image/color"
    "github.com/nfnt/resize"
)

func AddMap(i image.Image) image.Image{
    mapFile, err := os.Open("./res/map.png")
    if err != nil { panic(err) }
    mapImg, _, err := image.Decode(mapFile)
    if err != nil { panic(err) }
    mask := image.NewUniform(color.Alpha{128})

    canvas := image.NewRGBA(mapImg.Bounds())
    draw.Draw(canvas, canvas.Bounds(), mapImg, image.Point{0, 0}, draw.Src)

    draw.DrawMask(canvas, canvas.Bounds(), i, image.Point{0, 0}, mask, image.Point{0, 0}, draw.Over)

    return canvas
}

func ResizeToMap(i image.Image) image.Image{
    return resize.Resize(1491, 836, i, resize.Lanczos3)
}

func SaveToRecords(i image.Image, fileName string) {
    checkAndCreateDir("./records")

    file, err := os.Create("./records/" + fileName)
    if err != nil { panic(err) }
    defer file.Close()

    png.Encode(file, i)
}

func AlphaPercent(i image.Image) float64 {
    var opaquePixs int
    width, height := 1491, 836

    for x := 0; x <= width; x++ {
        for y:= 0; y <= height; y++{
            _, _, _, alphaValue := i.At(x, y).RGBA()
            if alphaValue != 0 {
                opaquePixs++
            }
        }
    }
    return float64(opaquePixs) / float64(width * height)
}

func AlphaPercentMask(i image.Image) float64 {
    var opaquePixs int
    width, height := 1491, 836
    maskFile, err := os.Open("./res/mask.png")
    if err != nil { panic (err) }
    mask, _, err := image.Decode(maskFile)
    if err != nil { panic (err) }

    for x := 0; x <= width; x++ {
        for y := 0; y <= height; y++ {
            _, _, _, maskAlphaValue := mask.At(x, y).RGBA()
            _, _, _, alphaValue := i.At(x, y).RGBA()
            if maskAlphaValue != 0 && alphaValue != 0 {
                opaquePixs++
            }
        }
    }
    return float64(opaquePixs) / float64(width * height)
}
