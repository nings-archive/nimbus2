package lib

import (
    "io/ioutil"
    "encoding/json"
)

type keys struct {
    ApiKey string
}

func DecodeJson(name string) keys {
    _keys := keys{}
    file, err := ioutil.ReadFile(name)
    if err != nil { panic(err) }

    err = json.Unmarshal(file, &_keys)
    if err != nil { panic(err) }
    return _keys
}
