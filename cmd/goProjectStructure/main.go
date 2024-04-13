package main

import (
	"fmt"
	"github.com/sjones1997/goProjectStructure/internal/routes"
	"net/http"
)

func main() {

	route := routes.NewRouter()

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Println("Listening on ", addr)

	err := http.ListenAndServe(addr, route)

	if err != nil {
		panic(err)
	}

}
