package main

import (
	"embed"
	"fmt"
	_ "friends_ranking/bootstrap"
	"friends_ranking/router"
)

var yamlFile embed.FS

func main() {
	fmt.Println("friends_ranking start...")
	r := router.InitRouter()
	r.Run(":8809")
}
