package lib

import (
    "os"
)

func CheckAndCreate(path string) {
    exists, err := DirExists(path)
    if err == nil {
        if !exists {
            os.Mkdir(path, 0755)
        }
    } else {
        panic(err)
    }
}

func DirExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil}
    return true, err
}
