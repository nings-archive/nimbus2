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
    mapImgNrgba, _, err := image.Decode(mapFile)
    mapImg := image.NewRGBA(mapImgNrgba.Bounds())
    draw.Draw(mapImg, mapImg.Bounds(), mapImgNrgba, image.Point{0, 0}, draw.Src)
    if err != nil { panic(err) }
    
    mask := image.NewUniform(color.Alpha16{32767})

    canvas := image.NewRGBA(mapImg.Bounds())
    draw.Draw(canvas, canvas.Bounds(), mapImg, image.Point{0, 0}, draw.Src)

    draw.DrawMask(canvas, canvas.Bounds(), i, image.Point{0, 0}, mask, image.Point{0, 0}, draw.Over)

    return canvas
}

func ResizeToMap(i image.Image) image.Image{
    return resize.Resize(1491, 836, i, resize.Bilinear)
}

func DebugResizes(i image.Image) {
    SaveToRecords(resize.Resize(1491, 836, i, resize.NearestNeighbor), "nearestNeighbor.png")
    SaveToRecords(resize.Resize(1491, 836, i, resize.Bilinear), "bilinear.png")
    SaveToRecords(resize.Resize(1491, 836, i, resize.Bicubic), "bicubic.png")
    SaveToRecords(resize.Resize(1491, 836, i, resize.MitchellNetravali), "mitchellNetravali.png")
    SaveToRecords(resize.Resize(1491, 836, i, resize.Lanczos2), "lanczos2.png")
    SaveToRecords(resize.Resize(1491, 836, i, resize.Lanczos3), "lanczos3.png")
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
