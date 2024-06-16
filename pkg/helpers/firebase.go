package helpers

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/auth"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	"github.com/gin-gonic/gin"
)

type UserTest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Picture  string `json:"picture"`
}

// getAuthClient returns an instance of the Firebase Authentication client.
// It takes a Gin context as a parameter and returns a pointer to the auth.Client.
// If an error occurs during the creation of the client, it logs the error and returns nil.
func getAuthClient(c *gin.Context) *auth.Client {
	authClient, err := global.AdminSdk.Auth(c.Request.Context())
	if err != nil {
		errMsg := fmt.Errorf("error creating user: %v", err)
		log.Fatalf(errMsg.Error())
		return nil
	}

	return authClient
}

// GetUserIDToken retrieves the user ID token and returns a SocialResponse object containing user information.
// It takes a gin.Context object and a uid string as parameters.
// If the user ID token cannot be retrieved or the authClient is nil, it returns nil.
// Otherwise, it returns a SocialResponse object with the user's full name, email, and picture.
func GetUserIDToken(c *gin.Context, uid string) *models.SocialResponse {
	authClient := getAuthClient(c)

	userRecord, err := authClient.GetUser(context.Background(), uid)
	if err != nil || authClient == nil {
		return nil
	}

	return &models.SocialResponse{
		Fullname: userRecord.DisplayName,
		Email:    userRecord.Email,
		Picture:  userRecord.PhotoURL,
	}
}

// GetUserUIDByEmail retrieves a user's UID by their email address.
// It takes a gin.Context and the user's email address as input.
// It returns the UID of the user and any error encountered during the retrieval.
func GetUserUIDByEmail(c *gin.Context, email string) (string, error) {
	authClient := getAuthClient(c)

	userRecord, err := authClient.GetUserByEmail(context.Background(), email)
	if err != nil {
		return "", fmt.Errorf("error retrieving user by email: %v", err)
	}
	fmt.Printf("Successfully retrieved user UID: %v for email: %v\n", userRecord.UID, email)
	return userRecord.UID, nil
}

// createUser creates a new user in Firebase Authentication with the provided email and password.
// It returns the created user record or an error if the user creation fails.
func CreateUser(c *gin.Context, email, password string) (*auth.UserRecord, error) {
	authClient := getAuthClient(c)

	params := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(true).
		Password(password).
		DisplayName(email).
		Disabled(false)

	u, err := authClient.CreateUser(context.Background(), params)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}
	fmt.Printf("Successfully created user: %v\n", u.UID)
	return u, nil
}

// UpdateUserEmail updates the email address of a user in Firebase Authentication.
// It takes a gin.Context, user ID (uid), and the new email address as input.
// It returns the updated UserRecord and any error encountered during the update.
func UpdateUserEmail(c *gin.Context, uid, newEmail string) (*auth.UserRecord, error) {
	authClient := getAuthClient(c)

	params := (&auth.UserToUpdate{}).
		Email(newEmail).
		EmailVerified(true)

	u, err := authClient.UpdateUser(context.Background(), uid, params)
	if err != nil {
		return nil, fmt.Errorf("error updating user email: %v", err)
	}
	log.Printf("Successfully updated user email: %v\n", u.Email)
	return u, nil
}

// DeleteUser deletes a user from Firebase Authentication using the provided user ID.
// It returns an error if there was a problem deleting the user.
func DeleteUser(c *gin.Context, uid string) error {
	authClient := getAuthClient(c)

	err := authClient.DeleteUser(context.Background(), uid)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	fmt.Printf("Successfully deleted user: %v\n", uid)
	return nil
}
