package otus_social

import (
	"log"
	"otus_social/controllers"
)

func main() {
	server := controllers.CreateServer()

	err := server.Run()
	if err != nil {
		log.Fatalf("Failed to start server: %q", err)
		return
	}
}
