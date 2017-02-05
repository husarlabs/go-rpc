package main

import (
    "fmt"
    rpc "github.com/husarlabs/go-rpc"
)

func main() {
    client, err := rpc.NewClient(nil, "http://localhost:5050")
    if err != nil {
        fmt.Println(err)
        return
    }
    m := make(map[string]interface{})
    m["id"] = 1
    r := make(map[string]interface{})
    err = client.Call("method", &m, &r)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(r)
}