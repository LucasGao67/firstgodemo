package router

import (
	"github.com/gin-gonic/gin"
	"github.com/LucasGao67/firstgodemo/router/middleware"
	"net/http"
	"github.com/LucasGao67/firstgodemo/handler/sd"
	"github.com/LucasGao67/firstgodemo/handler/user"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route")
	})

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	u := g.Group("/v1/user")
	{
		u.POST("",user.Create)
	}

	return g
}
