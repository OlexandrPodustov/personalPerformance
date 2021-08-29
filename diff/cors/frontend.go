package main

import (
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./frontend", false)))

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
