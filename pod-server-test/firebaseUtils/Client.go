package firebaseutils

import (
	"context"
	"fmt"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	utils "github.com/shark121/pods-test/utils"
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
	Client *firestore.Client
	Ctx    context.Context
}

func (app *Instance) Add(data map[string]any) {
	collection := app.Client.Doc("pods/9vy420")

	collection.Set(app.Ctx, data)

	print("data added successfully")
}

func CreateAppInstance() (Instance, error) {
	var ctx context.Context = context.Background()

	options := option.WithCredentialsFile("pods-rideshare-firebase-adminsdk-fbsvc-7098514035.json")

	config := &firebase.Config{
		ProjectID: "pods-rideshare",
	}

	app, err := firebase.NewApp(ctx, config, options)

	if err != nil {
		fmt.Println("error creating app instance", err)
	}

	db, err := app.Firestore(ctx)

	utils.HasErr(err)

	return Instance{db, ctx}, err
}
