package main

import (
	_ "lottery_annual/bootstrap"
	"lottery_annual/config/variable"
	"lottery_annual/router"
)

func main() {
	r := router.InitRouter()
	_ = r.Run(variable.YamlConfig.GetString("Application.Port"))

}
