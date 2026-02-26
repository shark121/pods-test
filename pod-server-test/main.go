package main

import (
<<<<<<< HEAD
	"net/http"

	firebaseutils "github.com/shark121/pods-test/firebaseUtils"
	"github.com/shark121/pods-test/helpers"
	"github.com/shark121/pods-test/utils"
)

func main() {

	// podAndRides := map[string]any{"randomRides": randomRides, "pod": pod, "ranked": rankRidesByProximityToPod(randomRides, pod)}
	appInstance, err := firebaseutils.CreateAppInstance()

	utils.HasErr(err)

	helpers.UseHandler(appInstance)

	appInstance.Add(map[string]any{"hello": "world"})

	print("server started running")

	http.ListenAndServe(":5000", nil)

=======
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
>>>>>>> d39e2fc54b334054a6b30ba699952213009a9465
}
