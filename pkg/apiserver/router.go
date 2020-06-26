package apiserver

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (sv *server) BuildRouter() *gin.Engine {
	router := gin.New()
	ginpprof.Wrap(router)

	// 项目
	router.GET("/project/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "get success")
	})
	router.POST("/project", func(c *gin.Context) {
		c.String(http.StatusOK, "add success")
	})
	router.PUT("/project", func(c *gin.Context) {
		c.String(http.StatusOK, "edit success")
	})
	router.DELETE("/project/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "delete success")
	})

	// 分组
	router.GET("/group/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "get success")
	})
	router.POST("/group", func(c *gin.Context) {
		c.String(http.StatusOK, "add success")
	})
	router.PUT("/group", func(c *gin.Context) {
		c.String(http.StatusOK, "edit success")
	})
	router.DELETE("/group/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "delete success")
	})

	// API
	router.GET("/api/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "get success")
	})
	router.POST("/api", func(c *gin.Context) {
		c.String(http.StatusOK, "add success")
	})
	router.PUT("/api", func(c *gin.Context) {
		c.String(http.StatusOK, "edit success")
	})
	router.DELETE("/api/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "delete success")
	})

	router.GET("/api/:id/export", func(c *gin.Context) {
		c.String(http.StatusOK, "export success default yaml extention")
	})

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusOK, "no router")
	})

	router.NoMethod(func(c *gin.Context) {
		c.String(http.StatusOK, "unsupported method")
	})
	return router
}
