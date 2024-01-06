package reserveca

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func (c *Client) CheckCampsite(reqData CampsiteRequest) (*CampsiteResponse, error) {
    req, err := c.newRequest("POST", c.BaseURL, reqData)
    if err != nil {
        return nil, err
    }

    resp, err := c.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API request failed with status code %d", resp.StatusCode)
    }

    var response CampsiteResponse
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        return nil, err
    }

    return &response, nil
}
