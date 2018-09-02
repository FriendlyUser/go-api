package main

import "os"


func main() {
	app := App{}
	app.Initialize(
		os.Getenv("connectionString"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "5432"
	}
	app.Run(":" + port)
}