package main

import (
	"embed"
	_ "friends_ranking/bootstrap"
	"friends_ranking/router"
)

var yamlFile embed.FS

func main() {
	r := router.InitRouter()
	r.Run(":8809")
}
