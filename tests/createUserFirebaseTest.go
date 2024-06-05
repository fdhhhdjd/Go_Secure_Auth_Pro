package tests

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/auth"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/pkg/helpers"
	"github.com/gin-gonic/gin"
)

type UserTest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Picture  string `json:"picture"`
}

func CreateUserFirebaseTest(c *gin.Context) {
	authClient, err := global.AdminSdk.Auth(c.Request.Context())
	if err != nil {
		errMsg := fmt.Errorf("error creating user: %v", err)
		log.Fatalf(errMsg.Error())
	}

	// Create a new user
	email := helpers.RandomEmail()
	password := helpers.RandomPassword()
	u, err := createUser(authClient, email, password)
	if err != nil {
		errMsg := fmt.Errorf("error creating user: %v", err)
		log.Fatalf(errMsg.Error())
	}

	// Get the ID Token for the user
	userRecord := getUserIDToken(authClient, u.UID)
	if userRecord != nil {
		fmt.Printf("ID userRecord: %s\n", userRecord)
	} else {
		log.Fatalf("error getting ID userRecord")
	}

}

func createUser(authClient *auth.Client, email, password string) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(false).
		Password(password).
		DisplayName("Example User").
		Disabled(false)

	u, err := authClient.CreateUser(context.Background(), params)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}
	fmt.Printf("Successfully created user: %v\n", u.UID)
	return u, nil
}

func getUserIDToken(authClient *auth.Client, uid string) *UserTest {
	userRecord, err := authClient.GetUser(context.Background(), uid)
	if err != nil {
		return nil
	}

	return &UserTest{
		Fullname: userRecord.DisplayName,
		Email:    userRecord.Email,
		Picture:  userRecord.PhotoURL,
	}
}
