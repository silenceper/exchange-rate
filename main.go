package main

import (
	"runtime"

	"github.com/gin-gonic/gin"
)

var cfg *config

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	cfg = getConfig()

	r := gin.Default()
	registerRouter(r)
	r.Run(cfg.Listen)
}
