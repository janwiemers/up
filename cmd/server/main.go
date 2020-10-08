package main

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/janwiemers/up/handler"
	"github.com/janwiemers/up/helper"
	"github.com/janwiemers/up/models"
	"github.com/janwiemers/up/monitors"
	"github.com/janwiemers/up/websockets"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func init() {
	helper.InitViperConfig()
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(log.InfoLevel)
}

func loadAndInitialzeConfigs() {
	monitorConfigs := []models.Application{}
	err := yaml.Unmarshal(helper.ReadFile(viper.GetString("MONITOR_FILE_PATH")), &monitorConfigs)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	monitors.InitAllMonitors(monitorConfigs)
	monitors.Cleanup(monitorConfigs)
}

func main() {
	loadAndInitialzeConfigs()
	go helper.Cleanup()
	websockets.HubInstance = websockets.NewHub()
	go websockets.HubInstance.Run()
	gin.DisableConsoleColor()
	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	p := ginprom.New(
		ginprom.Engine(r),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)

	r.Use(p.Instrument())
	r.Use(cors.New(config))
	r.Use(gin.Recovery())
	handler.SetupRouter(r)
	r.Run(":" + viper.GetString("PORT"))
}
