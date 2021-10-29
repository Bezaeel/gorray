package api

import (
	"fmt"

	"github.com/bezaeel/gorray/internal/api/router"
	"github.com/bezaeel/gorray/internal/pkg/config"
)

func setConfiguration(configPath string){
	config.Setup(configPath)

}

func Run(configPath string) {
	if configPath == "" {
		configPath = "../../config/"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()
	fmt.Println("Go API REST Running on port " + conf.Server.Port)
	fmt.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}