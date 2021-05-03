package infrastructure

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func NewFirebaseApp() *firebase.App {
	opt := option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_SERVICE_ACCOUNT_JSON")))
	// app, err := firebase.NewApp(ctx, nil, opt)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err.Error())
	}
	return app
}

func NewAuthClient(app *firebase.App) *auth.Client {
	// Access auth service from the default app
	client, err := app.Auth(context.Background())
	if err != nil {
		panic(err.Error())
	}
	return client
}
