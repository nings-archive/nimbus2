package lib

import (
    "os"
    "strconv"
    "encoding/csv"
)

func AppendCSV(fileName string, alphaPercent float64) {
    checkAndCreateCSV()
    alphaString := strconv.FormatFloat(alphaPercent, 'f', 10, 64)
    newEntry := []string{fileName, alphaString}

    file, err := os.OpenFile("./records/record.csv", os.O_RDWR|os.O_APPEND, 0755)
    if err != nil { panic(err) }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    err = writer.Write(newEntry)
    if err != nil { panic(err) }
}

func ReadCSV() [][]string {
    checkAndCreateCSV()
    file, err := os.Open("./records/record.csv")
    if err != nil { panic(err) }
    defer file.Close()

    reader := csv.NewReader(file)

    entries, err := reader.ReadAll()
    if err != nil { panic(err) }
    return entries
}

func GetLastEntry() []string {
    checkAndCreateCSV()
    csvEntries := ReadCSV()
    return csvEntries[len(csvEntries)-1]
}

func checkAndCreateCSV() {
    checkAndCreateDir("./records")
    checkAndCreateFile("./records/record.csv")
}

