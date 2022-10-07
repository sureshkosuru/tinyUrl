package main

import (
  "bytes"
  "encoding/json"
  "errors"
  "fmt"
  "io"
  "net/http"
  "time"
)

func fireGetRequest(req getUrlInfoReq, resp *getUrlInfoResp)(error) {
  respBytes, err := fireHttpRequest("GET", req)
  if err != nil {
    return err
  }
  err = json.Unmarshal(respBytes, &resp)
  if err != nil {
    return err
  }
  return nil
}

func firePutRequest(req putUrlReq, resp *putUrlResp)(error) {
  respBytes, err := fireHttpRequest("PUT", req)
  if err != nil {
    return err
  }
  err = json.Unmarshal(respBytes, resp)
  if err != nil {
    return err
  }
  return nil
}

func fireDeleteRequest(req deleteUrlReq, resp *deleteUrlResp)(error) {
  respBytes, err := fireHttpRequest("DELETE", req)
  if err != nil {
    return err
  }
  err = json.Unmarshal(respBytes, resp)
  if err != nil {
    return err
  }
  return nil
}

func fireHttpRequest(reqType string, req interface{})([]byte, error) {
  dataReq, err := json.Marshal(req)
  if err != nil {
    return nil,err
  }

  c := http.Client{Timeout: time.Duration(5) * time.Second}
  serverUrl := "http://" + *serverPort + *endPoint
  httpReq, err := http.NewRequest(reqType, serverUrl, bytes.NewReader(dataReq))
  if err != nil {
    fmt.Printf("error %s", err)
    return nil, err
  }

  httpReq.Header.Add("Accept", `application/json`)

  dataResp, err := c.Do(httpReq)
  if err != nil {
    fmt.Printf("Error %s", err)
    return nil, err
  }

  if err != nil {
    return nil, err
  }

  if dataResp.StatusCode != http.StatusOK {
    return nil, errors.New("Http request failed.")
  }
  defer dataResp.Body.Close()
  respBytes, err := io.ReadAll(dataResp.Body)
  if err != nil {
    return nil, err
  }
  return respBytes, nil
}
