package reserveca

import (
    "bytes"
    "encoding/json"
    "net/http"
)

type Client struct {
    httpClient *http.Client
    BaseURL    string
}

func NewClient() *Client {
    return &Client{
        httpClient: &http.Client{},
        BaseURL:    "https://calirdr.usedirect.com/rdr/rdr/search/grid",
    }
}

func (c *Client) newRequest(method, url string, body interface{}) (*http.Request, error) {
    buf := new(bytes.Buffer)
    if body != nil {
        err := json.NewEncoder(buf).Encode(body)
        if err != nil {
            return nil, err
        }
    }

    req, err := http.NewRequest(method, url, buf)
    if err != nil {
        return nil, err
    }

    req.Header.Add("Content-Type", "application/json")
    return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
    return c.httpClient.Do(req)
}
