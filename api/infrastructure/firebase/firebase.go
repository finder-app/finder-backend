package firebase

import (
	"context"
	"log"

	"api/infrastructure/env"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type FirebaseApp interface {
	Auth(ctx context.Context) (*auth.Client, error)
}

type FirebaseClient interface {
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
}

func NewFirebaseApp() FirebaseApp {
	// Initialize default app
	opt := option.WithCredentialsJSON([]byte(env.GOOGLE_SERVICE_ACCOUNT_JSON))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return app
}

func NewClient(app FirebaseApp) FirebaseClient {
	// Access auth service from the default app
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}
	return client
}
