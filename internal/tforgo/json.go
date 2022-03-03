package tforgo

import (
    "encoding/json"
    "fmt"
    "log"
)

type Message struct {
    Name string
    Body string
    Time int64
}

func TjsonTest() {
    m := Message{"Alice", "Hello", 1294706395881547000}
    b, err := json.Marshal(m)

    if err != nil {
        fmt.Println("json err:", err)
    }

    fmt.Printf("b: %s\n", string(b))

    raw_json := "{\"Name\":\"Bob\",\"Body\":\"Hello world\",\"Time\":1294706395881547111}"
    str := []byte(raw_json)

    new_msg := Message{}
    err_1 := json.Unmarshal(str, &new_msg)

    if err_1 != nil {
        fmt.Println("json Unmarshal err:", err)
    }

    log.Printf("new_msg: %+v\n", new_msg)
}
