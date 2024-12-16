package main

import (
	"fmt"
	"github.com/hudayberdipolat/go-web-temp-crud/models"
	"github.com/hudayberdipolat/go-web-temp-crud/routes"
	"net/http"
)

func main() {

	models.Post{}.Migrate()
	fmt.Println("Server : https://localhost:8080")
	http.ListenAndServe(":8080", routes.Routes())
}
