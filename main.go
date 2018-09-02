package main

import "os"

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func main() {
	app := App{}
	app.Initialize(
		genENV("APP_DB_USERNAME","circleci"),
		os.Getenv("APP_DB_PASSWORD",""),
		os.Getenv("APP_DB_NAME","circle_test"),
		os.Getenv("APP_DB_HOST","postgres"),
		"5432",
		"require")

	port := os.Getenv("PORT")
	if port == "" {
		port = "5432"
	}
	app.Run(":" + port)
}
