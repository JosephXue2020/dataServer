package main

import (
	"fmt"

	"goweb/dataServer/dao"
	"goweb/dataServer/router"
)

func main() {
	fmt.Println("Data server website is starting...")

	err := dao.InitSqlite()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()

	r := router.SetupRouter()
	r.Run(":8000")
}
