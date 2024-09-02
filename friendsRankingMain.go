package main

import (
	"fmt"
	"friends_ranking/config/variable"
	"friends_ranking/config/yamlConfig"
	"friends_ranking/router"
)

func main() {
	fmt.Println("hello go !")
	yamlConfig.CreateYamlFactory("")
	applicationName := variable.YamlConfig.GetString("application.name")
	fmt.Println("applicationaName: ", applicationName)
	r := router.InitRouter()
	r.Run(":8809")
}
