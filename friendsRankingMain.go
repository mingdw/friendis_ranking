package main

import (
	"fmt"
	"friends_ranking/router"
)

func main() {
	fmt.Println("hello go !")
	r := router.InitRouter()
	r.Run(":8809")
}
