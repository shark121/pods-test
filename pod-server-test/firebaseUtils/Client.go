package firebaseutils

import (
	"context"
	"fmt"

	"github.com/pod-server-test/utils"

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
	DB  *firestore.Client
	Ctx context.Context
}

func (app *Instance) Add(data map[string]any) {
	// collection := app.DB.Doc("students/names")

	// collection.Set(app.Ctx, data)

	// print("data added successfully")
}

func CreateAppInstance() (Instance, error) {
	var ctx context.Context = context.Background()
	cfg := utils.LoadConfig()

	var options option.ClientOption
	if cfg.FirebaseKeyPath != "" {
		options = option.WithCredentialsFile(cfg.FirebaseKeyPath)
	} else {
		// Used deployed service account logic if path is empty
		options = option.WithCredentialsFile("")
	}

	app, err := firebase.NewApp(ctx, nil, options)

	if err != nil {
		return Instance{}, fmt.Errorf("error creating app instance: %w", err)
	}

	db, err := app.Firestore(ctx)

	if err != nil {
		return Instance{}, fmt.Errorf("error creating firestore client: %w", err)
	}

	return Instance{DB: db, Ctx: ctx}, nil
}
