package main

import (
	"log"

	"github.com/Sri2103/services/internal/api-gateway/dependency"
	http_transport "github.com/Sri2103/services/internal/api-gateway/transports/http"
)

func main() {
	dep, err := dependency.NewDependency()
	if err != nil {
		log.Fatal(err)
	}
	defer dep.ProductConn.Close()
	defer dep.CartConn.Close()
	defer dep.OrderConn.Close()
	err = http_transport.StartHttpServer(dep)
	if err != nil {
		log.Println("Error starting HTTP server: ", err.Error())
		dep.Logger.Fatal(err.Error())
	}

}
