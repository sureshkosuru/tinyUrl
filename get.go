package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

type getUrlInfoReq struct {
  ReqType     string    `json:"type,omitempty"`
  Url         []string  `json:"url,omitempty"`
  ShortendUrl []string  `json:"shortendurl,omitempty"`
}

type getUrlInfoResp struct {
  FoundUrls            []urlSet `json:"foundurls,omitempty"`
  NotFoundUrls         []urlSet `json:"notfoundurls,omitempty"`
  FoundShotenedUrls    []urlSet `json:"foundshortenedurls,omitempty"`
  NotFoundShotenedUrls []urlSet `json:"notfoundshortenedurls,omitempty"`
}

func getUrl(c *gin.Context) {
  var req getUrlInfoReq
  var resp getUrlInfoResp
  if err := c.BindJSON(&req); err != nil {
    return
  }

  var res urlSet
  res.Status = "Success"
  resp.FoundUrls = append(resp.FoundUrls, res)
  c.IndentedJSON(http.StatusOK, resp)
}
