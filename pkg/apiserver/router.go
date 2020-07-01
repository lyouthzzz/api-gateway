package apiserver

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/lyouthzzz/poseidon-api-gateway/pkg/log"
	"github.com/lyouthzzz/poseidon-api-gateway/pkg/middleware"
	"net/http"
)

func (sv *server) BuildRouter() *gin.Engine {
	router := gin.New()
	ginpprof.Wrap(router)

	router.Use(middleware.RequestId(), middleware.Logger(log.DefaultConf))

	apiRouter := router.Group("/api")
	// 项目
	apiRouter.GET("/project/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "get success")
	})
	apiRouter.POST("/project", func(c *gin.Context) {
		c.String(http.StatusOK, "add success")
	})
	apiRouter.PUT("/project", func(c *gin.Context) {
		c.String(http.StatusOK, "edit success")
	})
	apiRouter.DELETE("/project/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "delete success")
	})

	// 分组
	apiRouter.GET("/group/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "get success")
	})
	apiRouter.POST("/group", func(c *gin.Context) {
		c.String(http.StatusOK, "add success")
	})
	apiRouter.PUT("/group", func(c *gin.Context) {
		c.String(http.StatusOK, "edit success")
	})
	apiRouter.DELETE("/group/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "delete success")
	})

	// API
	apiRouter.GET("/api/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "get success")
	})
	apiRouter.POST("/api", func(c *gin.Context) {
		c.String(http.StatusOK, "add success")
	})
	apiRouter.PUT("/api", func(c *gin.Context) {
		c.String(http.StatusOK, "edit success")
	})
	apiRouter.DELETE("/api/:id", func(c *gin.Context) {
		c.String(http.StatusOK, "delete success")
	})

	apiRouter.GET("/api/:id/export", func(c *gin.Context) {
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
