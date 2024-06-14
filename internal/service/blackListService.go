package service

import (
	"log"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/global"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/gin-gonic/gin"
)

// BlackListIP is a function that handles the blacklisting of IP addresses.
// It receives a gin.Context object and returns a pointer to a slice of models.BodyIpRequest.
// The function checks the validity of the request body, adds the IP addresses to the blacklist,
// and returns the updated request body.
//
// Swagger documentation for BlackListIP function
// @Summary Blacklist IP addresses
// @Description Handles the blacklisting of IP addresses
// @Tags Blacklist
// @Accept json
// @Produce json
// @Param body body models.BodyIpRequest true "List of IP addresses to blacklist"
// @Success 200 {object} models.BodyIpRequest
// @Failure 400 {object} response.ErrorResponse
// @Router /blacklist/ip [post]
func BlackListIP(c *gin.Context) *models.BodyIpRequest {
	var reqBody models.BodyIpRequest

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.BadRequestError(c)
		return nil
	}

	if len(reqBody.IP) > 0 {
		// Add IP addresses to the blacklist
		for _, ip := range reqBody.IP {
			err := global.Cache.SAdd(c, constants.BlackListIP, ip).Err()
			if err != nil {
				log.Println("Failed to add IP to blacklist:", err)
				response.InternalServerError(c)
				return nil
			}
		}

	} else {
		response.BadRequestError(c)
		return nil
	}

	return &reqBody
}
