package main

import (
	"fmt"

	"goweb/dataServer/routers"
)

func main() {
	fmt.Println("Data server website is starting...")
	r := routers.Init()
	r.Run(":8000")
}
