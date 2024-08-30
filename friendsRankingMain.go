package main

import (
	"fmt"
	"friends_ranking/router"
)

func main() {
	fmt.Println("hello go !")
	r := router.UserRouter()
	r.Run(":8809")
}
