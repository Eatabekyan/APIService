package api

import (
	"authentification/dbconnect"
	"fmt"

	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// ////////////////////////////////////////
// Admin Handlers //
// ////////////////////////////////////////
func InsertHandler(ctx *gin.Context) {
	user, err := GetUserFromRequest(ctx)
	if err != nil {
		return
	}
	err = dbconnect.InsertUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Could not add user")
		log.Warning(fmt.Sprintf("Aborted with status 400: Could not add user\nError: %v\n", err))
		return
	}
	ctx.JSON(http.StatusOK, "User is successfully created.")
	log.Info("User is successfully created.")
}

func DeleteUserHandler(ctx *gin.Context) {
	user, err := GetUserFromRequest(ctx)
	if err != nil {
		return
	}
	err = dbconnect.DeleteUser(user.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Could not delete user")
		log.Warning(fmt.Sprintf("Aborted with status 400: Could not delete user\nError: %v\n", err))
		return
	}
	ctx.JSON(http.StatusOK, "User is successfully deleted.")
	log.Info("User is successfully deleted.")
}

func UpdateServicesHandler(ctx *gin.Context) {
	user, err := GetUserFromRequest(ctx)
	if err != nil {
		return
	}
	err = dbconnect.AddUserService(user)
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Could not update user")
		log.Warning(fmt.Sprintf("Aborted with status 400: Could not update user\nError: %v\n", err))
		return
	}
	ctx.JSON(http.StatusOK, "User is successfully updated.")
	log.Info("User is successfully updated.")
}

func RemoveServicesHanndler(ctx *gin.Context) {
	user, err := GetUserFromRequest(ctx)
	if err != nil {
		return
	}
	err = dbconnect.RemoveUserServices(user)
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Could not remove user")
		log.Warning(fmt.Sprintf("Aborted with status 400: Could not remove user\nError: %v\n", err))
		return
	}
	ctx.JSON(http.StatusOK, "Services are successfully removed.")
	log.Info("Seervices are successfully removed.")
}

//////////////////////////////////////////
