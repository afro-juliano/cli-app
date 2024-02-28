package todo

import (
	"encoding/json"
	"fmt"
    //"io/ioutil"
)

type Item struct {
    Text string
}

func SaveItems(filename string, items []Item) error {
    b, err := json.Marshal(items)
    if err != nil {
        return err
    }
    fmt.Println(string(b))
    return nil
}
