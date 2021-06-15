package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/tetsu-tech/playlist-team-shonosuke/infra"
	"github.com/tetsu-tech/playlist-team-shonosuke/project/config"
)

var router *gin.Engine

func Run() {
	router = gin.New()
	config.InitConfig()
	infra.InitMysql(
		config.Config.Mysql.DBTCPHost,
		config.Config.Mysql.Port,
		config.Config.Mysql.User,
		config.Config.Mysql.Password,
		config.Config.Mysql.Database,
		config.Config.Mysql.InstanceConnectionName,
	)

	router.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)
	if config.Config.App.Env == "local" {
		router.Use(gin.Logger())
		gin.SetMode(gin.DebugMode)
	}

	initRouters()
	err := router.Run(":" + config.Config.App.Port)
	if err != nil {
		log.Println(err)
	}
}

func initRouters() {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})
}
