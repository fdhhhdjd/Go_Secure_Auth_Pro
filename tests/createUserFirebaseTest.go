package tests

import (
	"fmt"
	"log"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/pkg/helpers"
	"github.com/gin-gonic/gin"
)

type UserTest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Picture  string `json:"picture"`
}

// CreateAndGetUidTestFireBase is a function that creates a new user in Firebase
// and retrieves the ID token for the user.
func CreateAndGetUidTestFireBase(c *gin.Context) {
	// Create a new user
	email := helpers.RandomEmail()
	password := helpers.RandomPassword()
	u, err := helpers.CreateUser(c, email, password)
	if err != nil {
		errMsg := fmt.Errorf("error creating user: %v", err)
		log.Fatalf(errMsg.Error())
	}

	// Get the ID Token for the user
	userRecord := helpers.GetUserIDToken(c, u.UID)
	if userRecord != nil {
		fmt.Printf("ID userRecord: %s\n", userRecord)
	} else {
		log.Fatalf("error getting ID userRecord")
	}
}
