package main

import (
	_ "friends_ranking/bootstrap"
	"friends_ranking/config/variable"
	"friends_ranking/router"
)

func main() {
	r := router.InitRouter()
	_ = r.Run(variable.YamlConfig.GetString("Application.Port"))

}
