package firebaseutils

import (
	"context"
	"fmt"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	option "google.golang.org/api/option"
)

type App struct {
	authOverride     map[string]any
	dbURL            string
	projectID        string
	serviceAccountID string
	storageBucket    string
	opts             []option.ClientOption
}

type Instance struct {
	db  *firestore.Client
	ctx context.Context
}

func (app *Instance) Add(data map[string]any) {
	// collection := app.db.Doc("students/names")

	// collection.Set(app.ctx, data)

	// print("data added successfully")
}

func CreateAppInstance() (Instance, error) {
	var ctx context.Context = context.Background()

	options := option.WithCredentialsFile("../../pods-rideshare-firebase-adminsdk-fbsvc-5b4e19c35f.json")

	app, err := firebase.NewApp(ctx, nil, options)

	if err != nil {
		fmt.Println("error creating app instance", err)
	}

	db, err := app.Firestore(ctx)

	if err != nil {
		print(err)
	}

	return Instance{db, ctx}, err
}

func main() {

	myDBApp, err := CreateAppInstance()

	if err != nil {
		fmt.Println(err)
	}

	stud := map[string]any{"h": "3"}

	myDBApp.Add(stud)

}
