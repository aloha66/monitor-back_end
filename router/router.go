package router

import (
	"github.com/gin-gonic/gin"
	"monitor-back_end/handler/dict"
	"monitor-back_end/handler/sd"
	"monitor-back_end/handler/user"
	"monitor-back_end/middleware"
	"net/http"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...) // Todo
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	u := g.Group("/v1/user")
	{
		u.POST("/:username", user.Create)
	}

	v1 := g.Group("/api/v1")
	{
		dictRouter := v1.Group("/dict")
		{
			dictRouter.GET("getDictList", dict.GetDictList)
			dictRouter.POST("createDict", dict.CreateDict)
		}
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
