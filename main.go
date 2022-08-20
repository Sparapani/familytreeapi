package main

import (
	"github.com/Sparapani/familytreeapi/api"
	"github.com/Sparapani/familytreeapi/services"
)

func main() {
	services.InitialCharge()
	api.StartRouter()
}
