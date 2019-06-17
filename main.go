package main

import "os"

func main() {
	a := App{}
	a.Initialize(os.Getenv("TEST_DB_USERNAME"), os.Getenv("TEST_DB_PASSWORD"), os.Getenv("TEST_DB_NAME"))

	a.Run(":8080")
}
