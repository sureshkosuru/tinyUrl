package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

type putUrlReq struct {
  Url         []string  `json:"url,omitempty"`
}

type putUrlResp struct {
  PutSuccessUrls []urlSet `json:"putsuccessurls,omitempty"`
  PutFailedUrls  []urlSet `json:"putfailedurls,omitempty"`
}

func putUrl(c *gin.Context) {
  var req putUrlReq
  var resp putUrlResp

  if err := c.BindJSON(&req); err != nil {
    return
  }

  var res urlSet
  res.Status = "Success"
  resp.PutSuccessUrls = append(resp.PutSuccessUrls, res)
  c.IndentedJSON(http.StatusOK, resp)
}
