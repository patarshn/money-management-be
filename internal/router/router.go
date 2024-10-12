package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/patarshn/money-management-be/config"
)

func InitRouter(cfg *config.Config) (*gin.Engine, error) {
	fmt.Println(cfg.App.Env)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r, nil
}

func Serve(cfg *config.Config, r *gin.Engine) {
	addr := fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port)
	r.Run(addr)
}
