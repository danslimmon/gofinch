package gofinch

import (
    "fmt"
    "errors"
    "strings"
    "net/http"
)

const finchBaseUrl = "http://f.exosite.com"


type Dataport struct {
    Id string
}
func (dp *Dataport) Write(val string) error {
    resp, err := post(dp.Id, val)
    if resp.StatusCode >= 300 {
        return errors.New(fmt.Sprintf("Received %d response when writing data to '%s'",
                                      resp.StatusCode, dp.Id))
    }
    return err
}
func NewDataport(dataportId string) *Dataport {
    return &Dataport{
        Id: dataportId,
    }
}


func post(uri, val string) (*http.Response, error) {
    url := fmt.Sprintf("%s%s", finchBaseUrl, uri)
    return http.Post(url, "text/plain", strings.NewReader(val))
}
