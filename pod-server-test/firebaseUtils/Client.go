package firebaseutils

import (
	"context"
	"fmt"

	"github.com/pod-server-test/utils"

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
<<<<<<< HEAD
	Client *firestore.Client
	Ctx    context.Context
}

func (app *Instance) Add(data map[string]any) {
	collection := app.Client.Doc("pods/9vy420")

	collection.Set(app.Ctx, data)
=======
	DB  *firestore.Client
	Ctx context.Context
}

func (app *Instance) Add(data map[string]any) {
	// collection := app.DB.Doc("students/names")

	// collection.Set(app.Ctx, data)
>>>>>>> d39e2fc54b334054a6b30ba699952213009a9465

	print("data added successfully")
}

func CreateAppInstance() (Instance, error) {
	var ctx context.Context = context.Background()
	cfg := utils.LoadConfig()

<<<<<<< HEAD
	options := option.WithCredentialsFile("pods-rideshare-firebase-adminsdk-fbsvc-7098514035.json")
=======
	var options option.ClientOption
	if cfg.FirebaseKeyPath != "" {
		options = option.WithCredentialsFile(cfg.FirebaseKeyPath)
	} else {
		// Used deployed service account logic if path is empty
		options = option.WithCredentialsFile("")
	}
>>>>>>> d39e2fc54b334054a6b30ba699952213009a9465

	config := &firebase.Config{
		ProjectID: "pods-rideshare",
	}

	app, err := firebase.NewApp(ctx, config, options)

	if err != nil {
		return Instance{}, fmt.Errorf("error creating app instance: %w", err)
	}

	db, err := app.Firestore(ctx)

<<<<<<< HEAD
	utils.HasErr(err)

	return Instance{db, ctx}, err
=======
	if err != nil {
		return Instance{}, fmt.Errorf("error creating firestore client: %w", err)
	}

	return Instance{DB: db, Ctx: ctx}, nil
>>>>>>> d39e2fc54b334054a6b30ba699952213009a9465
}
