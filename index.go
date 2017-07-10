package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nestgo/utils"
	"github.com/silenceper/exchange-rate/yahoo"
)

func doIndex(c *gin.Context) {
	c.Redirect(302, "/html")
}

func doExchange(c *gin.Context) {
	fromVal := c.Query("from_val")
	rateFrom := c.Query("rate_from")
	rateTo := c.Query("rate_to")

	if fromVal == "" || rateFrom == "" || rateTo == "" {
		utils.RenderError(c, -1, "参数不能为空")
		return
	}

	fromVal64, err := strconv.ParseFloat(fromVal, 64)
	if err != nil {
		utils.RenderError(c, -1, fmt.Sprintf("格式转换出错,err=%v", err))
		return
	}
	rate, err := yahoo.Exchange(rateFrom, rateTo)
	if err != nil {
		utils.RenderError(c, -1, fmt.Sprintf("转换出错,err=%v", err))
		return
	}
	utils.RenderSuccess(c, fromVal64*rate)
}
