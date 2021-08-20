package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"t-rex_wrapper/trex_wrapper"
	"t-rex_wrapper/utils"
)

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type Fan struct {
	Percent float64 `form:"percent" json:"percent" xml:"percent"`
}

func StartPage(c *gin.Context) {
	for i, s := range utils.Cfg.HostTarget {
		fmt.Println(i, s)
	}
	c.String(200, "Success")
}

func SetFanPercent(c *gin.Context) {
	fan := Fan{}
	if err := c.ShouldBindJSON(&fan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	trex_wrapper.SetFanPercent(fan.Percent)
	c.JSON(http.StatusOK, gin.H{"status": "Succes"})
}
