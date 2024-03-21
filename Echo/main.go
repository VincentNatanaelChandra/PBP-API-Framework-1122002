package main

import (
	"echo/db"
	"echo/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
