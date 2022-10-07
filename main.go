package main

import (
  "flag"
  "fmt"
  "strings"

  "github.com/gin-gonic/gin"
)

var shortenerToUrl map[string]string
var urlToShortener map[string]string
var maxPreFetchUrlShortener *uint
var endPoint *string
var serverPort *string

type urlSet struct {
  Status      string `json:"status,omitempty"`
  Url         string  `json:"url,omitempty"`
  ShortendUrl string  `json:"shortendurl,omitempty"`
}

func main() {
  gin.SetMode(gin.ReleaseMode)
  if parseConfAndNeedToDisplayHelp() {
    displayHelp()
  }

  fmt.Printf("\nshortAheadCount: %d\nendPoint: %s\nserverWithPort: %s\n\n", *maxPreFetchUrlShortener, *endPoint, *serverPort)
  go testMain()
  router := gin.Default()
  router.GET(*endPoint, getUrl)
  router.PUT(*endPoint, putUrl)
  router.DELETE(*endPoint, deleteUrl)
  fmt.Printf("\n\nStarting http server \n\n")

  router.Run(*serverPort)
}

func parseConfAndNeedToDisplayHelp() bool {
        shortOpt := flag.Bool("h", false, "display help")
        longOpt := flag.Bool("help", false, "display help")
        maxPreFetchUrlShortener = flag.Uint("shortAheadCount", 5, "Reserve short urls")
        endPoint = flag.String("endPoint", "/urlshortener", "Endpoint for http connection")
        serverPort = flag.String("serverWithPort", "10.163.164.135:8088", "Server details")
        flag.Parse()
        return *shortOpt || *longOpt
}

func displayHelp()  {
        help := strings.Builder{}
        help.WriteString("url-shortener [-h] [--endPoint=] [--shortAheadCount=] [--serverWithPort=]\n")
        help.WriteByte('\n')
        fmt.Println(help.String())
}
