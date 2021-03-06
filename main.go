package main

import "os"


func main() {
	app := App{}
	app.Initialize(
		os.Getenv("DATABASE_URL"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "5432"
	}
	app.Run(":" + port)
}