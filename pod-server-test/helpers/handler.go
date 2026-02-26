package helpers

import (
	"encoding/json"
	"fmt"
<<<<<<< HEAD
	"io"
	"net/http"

	firebaseutils "github.com/shark121/pods-test/firebaseUtils"
	podUtils "github.com/shark121/pods-test/podUtils"
	t "github.com/shark121/pods-test/types"

	check "github.com/shark121/pods-test/utils"
)

func CreateHandler(appInstance firebaseutils.Instance) func(res http.ResponseWriter, req *http.Request) {

	return func(res http.ResponseWriter, req *http.Request) {

		res.Header().Set("Content-Type", "application/json")                     // Set Content-Type
		res.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // Example Cache-Control
		res.Header().Set("Pragma", "no-cache")                                   // Example Pragma
		res.Header().Set("Expires", "0")
		res.Header().Set("Access-Control-Allow-Origin", "*")

		if req.Method != "POST" {
			print("wrong method")
		}

		requestBody := req.Body

		reqBytes, err := io.ReadAll(requestBody)

		check.HasErr(err)

		fmt.Println(reqBytes)

		reqBodyJson := t.RideObject{}

		json.Unmarshal(reqBytes, &reqBodyJson)

		fmt.Println(reqBodyJson)

		pod := podUtils.MatchRide(reqBodyJson)

		podUtils.HandleRideRequest(appInstance.Ctx, appInstance.Client, reqBodyJson)

		// response, err := json.Marshal(data)

		// if err != nil {
		// 	print(err)
		// }
		json.NewEncoder(res).Encode(pod)

	}
=======
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	firebaseutils "github.com/pod-server-test/firebaseUtils"
	t "github.com/pod-server-test/types"
)

func SetupRoutes(mux *http.ServeMux, dbApp firebaseutils.Instance) {
	mux.HandleFunc("/api/request-ride", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		if req.Method != http.MethodPost {
			http.Error(res, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
			return
		}

		var rideReq struct {
			Origin      t.Location `json:"origin"`
			Destination t.Location `json:"destination"`
			Capacity    int8       `json:"capacity"`
		}

		if err := json.NewDecoder(req.Body).Decode(&rideReq); err != nil {
			http.Error(res, `{"error": "Invalid request body"}`, http.StatusBadRequest)
			return
		}

		ride := t.RideObject{
			RideID:       uuid.New().String(),
			RideTime:     time.Now().Format(time.RFC3339),
			RideStatus:   "pending",
			Origin:       rideReq.Origin,
			Destination:  rideReq.Destination,
			RideCapacity: rideReq.Capacity,
			Direction:    0.0,
			CreatedAt:    time.Now(),
		}

		pod, err := HandleRideRequest(dbApp.Ctx, dbApp.DB, ride)
		if err != nil {
			http.Error(res, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(res).Encode(map[string]interface{}{
			"message": "Ride request processed",
			"rideId":  ride.RideID,
			"pod":     pod,
		})
	})

	mux.HandleFunc("/api/pod/extend", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		if req.Method != http.MethodPost {
			http.Error(res, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
			return
		}

		var podReq struct {
			PodID    string `json:"podId"`
			Capacity int8   `json:"capacity"`
		}

		if err := json.NewDecoder(req.Body).Decode(&podReq); err != nil {
			http.Error(res, `{"error": "Invalid request body"}`, http.StatusBadRequest)
			return
		}

		tierDoc := fmt.Sprintf("tier_%d", podReq.Capacity)
		_, err := dbApp.DB.Collection("pods").Doc(tierDoc).Collection("activePods").Doc(podReq.PodID).Update(dbApp.Ctx, []firestore.Update{
			{Path: "createdAt", Value: time.Now()},
		})

		if err != nil {
			http.Error(res, `{"error": "Failed to extend pod"}`, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(res).Encode(map[string]interface{}{
			"message": "Pod wait time extended by 30 minutes",
		})
	})
>>>>>>> d39e2fc54b334054a6b30ba699952213009a9465

	mux.HandleFunc("/api/pod/dispatch", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		if req.Method != http.MethodPost {
			http.Error(res, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
			return
		}

<<<<<<< HEAD
func UseHandler(appInstance firebaseutils.Instance) {
	getHandler := CreateHandler(appInstance)
	http.HandleFunc("/", getHandler)
=======
		var podReq struct {
			PodID    string `json:"podId"`
			Capacity int8   `json:"capacity"`
		}

		if err := json.NewDecoder(req.Body).Decode(&podReq); err != nil {
			http.Error(res, `{"error": "Invalid request body"}`, http.StatusBadRequest)
			return
		}

		tierDoc := fmt.Sprintf("tier_%d", podReq.Capacity)
		_, err := dbApp.DB.Collection("pods").Doc(tierDoc).Collection("activePods").Doc(podReq.PodID).Update(dbApp.Ctx, []firestore.Update{
			{Path: "podStatus", Value: "dispatched"},
		})

		if err != nil {
			http.Error(res, `{"error": "Failed to dispatch pod"}`, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(res).Encode(map[string]interface{}{
			"message": "Pod successfully dispatched",
		})
	})
>>>>>>> d39e2fc54b334054a6b30ba699952213009a9465
}
