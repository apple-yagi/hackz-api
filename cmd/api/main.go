package main

import (
	"hackz-api/pkg/db/pg"
	"hackz-api/pkg/routes"
)

func main() {
	defer pg.CloseConn()

	if err := routes.Router.Run(":8080"); err != nil {
		return
	}
}
