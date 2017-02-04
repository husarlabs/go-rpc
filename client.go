package rpc

import (
    "net/http"
    "net/url"
    "bytes"
    "github.com/gorilla/rpc/v2/json2"
)

type Client struct {
    server  *url.URL
    client  *http.Client
}

func NewClient(client *http.Client, urlp string) (*Client, error) {
    if client == nil {
        client = http.DefaultClient
    }

    server, err := url.Parse(urlp)

    if err != nil {
        return nil, err
    }

    if server.Scheme == "" {
        server.Scheme = "http"
    }
    if server.Scheme != "http" &&  server.Scheme != "https" {
        return nil, fmt.Errorf("Scheme is not supported")
    }

    return &Client {
        server: server,
        client: client,
    }, nil
}

func (c *Client) Call(method string, args interface{}, result interface{}) error {
    message, err := json2.EncodeClientRequest(method, args)
    if err != nil {
        return err
    }

    req, err := http.NewRequest("POST", c.server.String(), bytes.NewBuffer(message))
    if err != nil {
        return err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := c.client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    err = json2.DecodeClientResponse(resp.Body, &result)
    if err != nil {
        return err
    }
    return nil
}