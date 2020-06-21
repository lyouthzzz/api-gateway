package apiserver

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (sv *server) BuildRouter() *gin.Engine {
	engine := gin.New()
	ginpprof.Wrap(engine)

	// 项目
	engine.GET("/project/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "get success")
	})
	engine.POST("/project", func(c *gin.Context) {
		c.String(http.StatusOK, "add success")
	})
	engine.PUT("/project", func(c *gin.Context) {
		c.String(http.StatusOK, "edit success")
	})
	engine.DELETE("/project/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "delete success")
	})

	// 分组
	engine.GET("/group/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "get success")
	})
	engine.POST("/group", func(c *gin.Context) {
		c.String(http.StatusOK, "add success")
	})
	engine.PUT("/group", func(c *gin.Context) {
		c.String(http.StatusOK, "edit success")
	})
	engine.DELETE("/group/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "delete success")
	})

	// API
	engine.GET("/api/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "get success")
	})
	engine.POST("/api", func(c *gin.Context) {
		c.String(http.StatusOK, "add success")
	})
	engine.PUT("/api", func(c *gin.Context) {
		c.String(http.StatusOK, "edit success")
	})
	engine.DELETE("/api/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "delete success")
	})

	engine.GET("/api/:id/export", func(c *gin.Context) {
		c.String(http.StatusOK, "export success default yaml extention")
	})

	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusOK, "no router")
	})

	engine.NoMethod(func(c *gin.Context) {
		c.String(http.StatusOK, "unsupported method")
	})
	return engine
}
