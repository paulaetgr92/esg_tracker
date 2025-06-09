package main
import (
    "github.com/paulaetgr92/esg_tracker/models"
    "github.com/paulaetgr92/esg_tracker/routes"
	"log"
)

func main() {
	models.ConectarComBanco() 

	r := routes.SetupRouter()
	log.Fatal(r.Run(":8080")) 
}
