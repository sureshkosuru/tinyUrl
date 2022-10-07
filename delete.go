package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

type deleteUrlReq struct {
  ReqType       string `json:"reqtype,omitempty"`
  Url         []string `json:"url,omitempty"`
  ShortendUrl []string `json:"shortendurl,omitempty"`
}

type deleteUrlResp struct {
  DeleteSuccessUrls       []urlSet  `json:"deletesuccessurls,omitempty"`
  DeleteFailedUrls        []urlSet   `json:"deletefailedurls,omitempty"`
  DeleteSuccessShortends  []urlSet `json:"deletesuccessshortends,omitempty"`
  DeleteFailedShortends   []urlSet `json:"deletefailedshortends,omitempty"`
}

func deleteUrl(c *gin.Context) {
  var req deleteUrlReq
  var resp deleteUrlResp

  if err := c.BindJSON(&req); err != nil {
    return
  }

  var res urlSet
  res.Status = "Success"
  resp.DeleteSuccessUrls = append(resp.DeleteSuccessUrls, res)
  c.IndentedJSON(http.StatusOK, resp)
}
