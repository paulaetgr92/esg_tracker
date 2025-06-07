package main

import (
    "github.com/paulaetgr92/esg_tracker/controller"
    "github.com/paulaetgr92/esg_tracker/models"
    "github.com/paulaetgr92/esg_tracker/routes"
)


func main() {
	models.ConectarComBanco()
	r := routes.SetupRouter()
	r.Run(":8080")
}
