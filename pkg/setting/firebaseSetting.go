package pkg

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type UserUid struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Picture  string `json:"picture"`
}

// InitializeApp initializes the Firebase app and returns the app instance.
// It requires a service account key file path to authenticate with Firebase.
func InitializeApp() (*firebase.App, error) {
	opt := option.WithCredentialsFile("third_party/firebase/go-auth-pro-firebase-admin-sdk.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	fmt.Println("CONNECTED TO FIREBASE 🔥")
	return app, nil
}
