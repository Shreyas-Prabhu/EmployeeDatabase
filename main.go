package main

import (
	"fmt"

	"github.com/Shreyas-Prabhu/EmployeeDatabase/config"
	"github.com/Shreyas-Prabhu/EmployeeDatabase/router"
)

func main() {
	fmt.Println("GIN, MYSQL, Go")

	route := router.NewRouter()
	port := config.GetConfig("PORT")
	route.Run(port)
}
