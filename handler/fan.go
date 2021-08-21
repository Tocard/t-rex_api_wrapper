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
	FanConfig string `form:"fan_config" json:"fan_config" xml:"fan_config"`
}

func StartPage(c *gin.Context) {
	for i, s := range utils.Cfg.HostTarget {
		fmt.Println(i, s)
	}
	c.String(200, "Success")
}

//ManageFan is the handler to interact with all client trex fan
// @Summary Set value for remote fan
// @Description Post value to remote fan
// @Tags Fan
// @Success 200 {string} string
// @Failure 404 {string} string
// @Router / [post]
// @BasePath /managefan
func ManageFan(c *gin.Context) {
	fan := Fan{}
	if err := c.ShouldBindJSON(&fan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	trex_wrapper.ManageFan(fan.FanConfig)
	c.JSON(http.StatusOK, gin.H{"status": "Succes"})
}
