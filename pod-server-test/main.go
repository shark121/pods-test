package main

import (
	"log"
	"net/http"

	firebaseutils "github.com/pod-server-test/firebaseUtils"
	"github.com/pod-server-test/helpers"
	"github.com/pod-server-test/utils"
)

func main() {
	log.Println("Initializing application...")

	// 1. Load config
	utils.LoadConfig() // Ensure env variables are loaded

	// 2. Initialize Firebase
	dbApp, err := firebaseutils.CreateAppInstance()
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}
	defer dbApp.DB.Close()

	// 3. Setup Routes
	mux := http.NewServeMux()
	helpers.SetupRoutes(mux, dbApp)

	// 4. Start Server
	port := "5000"
	log.Printf("Server started running on port %s\n", port)

	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
