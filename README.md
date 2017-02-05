# go-rpc
A very simple helper library for calling JSON-RPC (v2) methods.

Install:
```bash
go get -u github.com/husarlabs/go-rpc
```

Import:
```go

import (
    "github.com/husarlabs/go-rpc"
)
```

Usage:

```go
args := make(map[string]interface{})
res := make(map[string]interface{})
args["id"] = 1

client, err := rpc.NewClient(nil, "http://localhost:5050")
if err != nil {
    return
}

err = client.Call("method", &args, &res)
if err != nil {
    return
}
fmt.Println(res)
```
