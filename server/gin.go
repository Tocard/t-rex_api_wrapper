package server

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"t-rex_wrapper/auth"
	"t-rex_wrapper/handler"
	"t-rex_wrapper/utils"
)

func engine() *gin.Engine {
	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))
	server.POST("/login", handler.Login)
	server.GET("/logout", handler.Logout)

	server.POST("/signup", handler.StartPage)
	private := server.Group("/trex")
	private.Use(AuthRequired)
	{
		private.GET("/me", me)
		private.GET("/status", status)
		private.POST("/managefan", handler.ManageFan)
	}
	return server
}

func me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(auth.Userkey)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(auth.Userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

func GoGinServer() {
	server := engine()
	server.Use(gin.Logger())
	if err := engine().Run(":" + fmt.Sprint(utils.Cfg.ApiPort)); err != nil {
		log.Fatal("Unable to start:", err)
	}
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
}
