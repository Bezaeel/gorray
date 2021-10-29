package router

import (
	"fmt"
	"io"
	"os"

	controller "github.com/bezaeel/gorray/internal/api/controllers/v1"
	"github.com/bezaeel/gorray/internal/api/middlewares"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	app := gin.New()

	// Logging to a file.
	f, _ := os.Create("log/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())

	// Routes
	// ================== Gorray Routes
	app.GET("/api/v1/echo", controller.Echo)
	app.GET("/api/v1/transpose", controller.Transpose)
	app.GET("/api/v1/flatten", controller.Flatten)
	app.GET("/api/v1/sum", controller.Sum)
	app.GET("/api/v1/multiply", controller.Multiply)
	return app
}