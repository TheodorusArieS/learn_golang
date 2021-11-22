// package main

// import (
// 	"net/http"
//     "github.com/gin-gonic/gin"
// )

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"title"`
// 	Price  float64 `json:"price"`
// }

// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 59.78},
// 	{ID: "2", Title: "Blue bagel", Artist: "elton Jon", Price: 61.21},
// 	{ID: "3", Title: "Blue waffle", Artist: "Mother Mode", Price: 12.12},
// }
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/pingpong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong ping",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}