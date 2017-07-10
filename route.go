package main

import "github.com/gin-gonic/gin"

func registerRouter(r *gin.Engine) {
	r.Static("/html", "./html")

	r.GET("/", doIndex)
	r.GET("/exchange", doExchange)
}
