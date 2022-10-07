package main

import (
  "errors"
  "fmt"
  "os"
  "time"
)

func testMain() {
  time.Sleep(10 * time.Second)
  fmt.Printf("\n\ntestMain: Started\n\n")
  err := testGet()
  if err != nil {
    fmt.Printf("testMain: %s\n", err)
    os.Exit(1)
  }

  err = testPut()
  if err != nil {
    fmt.Printf("testMain: %s\n", err)
    os.Exit(1)
  }

  err = testDelete()
  if err != nil {
    fmt.Printf("testMain: %s\n", err)
    os.Exit(1)
  }

  fmt.Printf("\n\ntestMain: Success\n\n")
  os.Exit(0)
}

func testGet() error {
  var req getUrlInfoReq
  var resp getUrlInfoResp
  req.ReqType = "4"
  req.Url = append(req.Url, "https://www.google.com")
  req.Url = append(req.Url, "https://www.amazon.com")
  req.ShortendUrl = append(req.ShortendUrl, "abcefgh")
  req.ShortendUrl = append(req.ShortendUrl, "ijklmno")
  err := fireGetRequest(req, &resp)
  if err != nil {
    return err
  }

  if len(resp.FoundUrls) != 1 {
	  return errors.New("testGet: Entries are not matching")
  } else {
    entity := resp.FoundUrls[0]
    if entity.Status != "Success" {
	    return errors.New("testGet: Status are not matching")
    }
  }

  return nil
}

func testPut() error {
  var req putUrlReq
  var resp putUrlResp
  req.Url = append(req.Url, "https://www.google.com")
  req.Url = append(req.Url, "https://www.amazon.com")
  err := firePutRequest(req, &resp)
  if err != nil {
    return err
  }

  if len(resp.PutSuccessUrls) != 1 {
	  return errors.New("testPut: Entries are not matching")
  } else {
    entity := resp.PutSuccessUrls[0]
    if entity.Status != "Success" {
	    return errors.New("testPut: Status are not matching")
    }
  }

  return nil
}

func testDelete() error {
  var req deleteUrlReq
  var resp deleteUrlResp
  req.ReqType = "4"
  req.Url = append(req.Url, "https://www.google.com")
  req.Url = append(req.Url, "https://www.amazon.com")
  req.ShortendUrl = append(req.ShortendUrl, "abcefgh")
  req.ShortendUrl = append(req.ShortendUrl, "ijklmno")
  err := fireDeleteRequest(req, &resp)
  if err != nil {
    return err
  }

  if len(resp.DeleteSuccessUrls) != 1 {
	  return errors.New("testDelete: Entries are not matching")
  } else {
    entity := resp.DeleteSuccessUrls[0]
    if entity.Status != "Success" {
      return errors.New("testDelete: Status are not matching")
    }
  }

  return nil
}
