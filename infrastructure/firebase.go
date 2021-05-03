package infrastructure

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func NewFirebaseApp(ctx context.Context) *firebase.App {
	opt := option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_SERVICE_ACCOUNT_JSON")))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		fmt.Errorf("error initializing app: %v", err)
	}
	return app
}

func NewAuthClient(app *firebase.App, ctx context.Context) *auth.Client {
	// Access auth service from the default app
	client, err := app.Auth(ctx)
	if err != nil {
		fmt.Errorf("error getting Auth client: %v", err)
	}
	return client
}
