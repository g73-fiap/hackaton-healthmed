package main

import "g73-fiap/hackaton-healthmed/internal/api"

func main() {
	port := "8080"

	api := api.NewApi(api.APIParams{})
	api.Run(":" + port)
}
