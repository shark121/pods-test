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

	utils.LoadConfig()

	dbApp, err := firebaseutils.CreateAppInstance()
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}
	defer dbApp.DB.Close()

	mux := http.NewServeMux()
	helpers.SetupRoutes(mux, dbApp)

	corsMux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		mux.ServeHTTP(w, r)
	})

	port := "5000"
	log.Printf("Server started running on port %s\n", port)

	err = http.ListenAndServe(":"+port, corsMux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
