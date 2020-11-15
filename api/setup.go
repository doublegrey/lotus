package api

import (
	"os"

	"github.com/doublegrey/lotus/api/app/apps"
	"github.com/doublegrey/lotus/api/app/logs"
	"github.com/doublegrey/lotus/api/service"
	"github.com/doublegrey/lotus/middlewares/logger"
	"github.com/doublegrey/lotus/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Setup returns configured gin router
func Setup() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	if !utils.Config.Server.Development {
		gin.SetMode(gin.ReleaseMode)
	}
	r.Use(logger.LogHandler())

	apiGroup := r.Group("api")
	serviceGroup := apiGroup.Group("service")
	appGroup := apiGroup.Group("app")

	serviceGroup.GET("settings", service.GetSettings)    // get lotus settings
	serviceGroup.PUT("settings", service.UpdateSettings) // update lotus settings
	serviceGroup.GET("health", service.Health)           // check lotus health

	appGroup.GET("", apps.GetAll)            // get list of registered apps and their settings
	appGroup.GET(":id", apps.Get)            // get app settings
	appGroup.POST("", apps.Create)           // create app
	appGroup.PUT(":id", apps.Update)         // update app settings
	appGroup.DELETE(":id", apps.Delete)      // delete app and its logs
	appGroup.GET(":id/logs", logs.Get)       // get app logs
	appGroup.POST(":id/logs", logs.Write)    // write message to app logs
	appGroup.DELETE(":id/logs", logs.Delete) // delete app logs

	return r
}
