package lib

import (
    "os"
)

func checkAndCreateFile(path string) {
    exists, err := isExists(path)
    if err == nil {
        if !exists {
            os.Create(path)
        }
    } else {
        panic(err)
    }
}

func checkAndCreateDir(path string) {
    exists, err := isExists(path)
    if err == nil {
        if !exists {
            os.Mkdir(path, 0755)
        }
    } else {
        panic(err)
    }
}

func isExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil}
    return true, err
}
