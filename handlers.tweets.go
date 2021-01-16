package main

import (
    // "net/http"
    // "strconv"
    "github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
  haiku := getHaiku(c)

  // Call the render function with the name of the template to render
  render(c, gin.H{
    "haiku":   haiku}, "index.html")

}
