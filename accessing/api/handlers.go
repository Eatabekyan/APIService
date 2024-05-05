package api

import (
	"accessing/dbconnect"
	"fmt"

	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// ////////////////////////////////////////
// Admin Handlers //
// ////////////////////////////////////////
func InsertHandler(log *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := getUserFromAdminRequest(log, ctx)
		if err != nil {
			return
		}
		err = dbconnect.InsertUser(user)
		if err != nil {
			ctx.AbortWithStatusJSON(400, fmt.Sprintf("Could not add user\nError: %v\n", err))
			log.Warning(fmt.Sprintf("Aborted with status 400: Could not add user\nError: %v\n", err))
			return
		}
		ctx.JSON(http.StatusOK, "User is successfully created.")
		log.Info("User is successfully created.")
	}
}

func DeleteUserHandler(log *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user, err := getUserFromAdminRequest(log, ctx)
		if err != nil {
			return
		}
		err = dbconnect.DeleteUser(user.Username)
		if err != nil {
			ctx.AbortWithStatusJSON(400, fmt.Sprintf("Could not delete user\nError: %v\n", err))
			log.Warning(fmt.Sprintf("Aborted with status 400: Could not delete user\nError: %v\n", err))
			return
		}
		ctx.JSON(http.StatusOK, "User is successfully deleted.")
		log.Info("User is successfully deleted.")
	}
}

func AddServicesHandler(log *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := getUserFromAdminRequest(log, ctx)
		if err != nil {
			return
		}
		err = dbconnect.AddUserService(user)
		if err != nil {
			ctx.AbortWithStatusJSON(400, fmt.Sprintf("Could not update user\nError: %v\n", err))
			log.Warning(fmt.Sprintf("Aborted with status 400: Could not update user\nError: %v\n", err))
			return
		}
		ctx.JSON(http.StatusOK, "User is successfully updated.")
		log.Info("User is successfully updated.")
	}
}

func RemoveServicesHanndler(log *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user, err := getUserFromAdminRequest(log, ctx)
		if err != nil {
			return
		}
		err = dbconnect.RemoveUserServices(user)
		if err != nil {
			ctx.AbortWithStatusJSON(400, fmt.Sprintf("Could not remove user\nError: %v\n", err))
			log.Warning(fmt.Sprintf("Aborted with status 400: Could not remove user\nError: %v\n", err))
			return
		}
		ctx.JSON(http.StatusOK, "Services are successfully removed.")
		log.Info("Services are successfully removed.")
	}
}

///////////////////////////////////////////
// Client Handlers //
///////////////////////////////////////////

func CheckUserAccessHandler(log *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqInfo, err := getClientRequestInfo(log, ctx)
		if err != nil {
			return
		}
		serviceName, err := checkToken(log, reqInfo.Token)
		if err != nil {
			ctx.AbortWithStatusJSON(400, "Access Denied")
			log.Warning(fmt.Sprintf("Access Denied, by token: %s\nError: %v", reqInfo.Token, err))
			return
		}
		user, err := dbconnect.GetUserByUsername(reqInfo.Username)
		if err != nil {
			ctx.AbortWithStatusJSON(400, fmt.Sprintf("Could not check user\nError: %v\n", err))
			log.Warning(fmt.Sprintf("Aborted with status 400: Could not check user\nError: %v\n", err))
			return
		}
		service, err := GetUserAccessToService(log, user, serviceName)
		if err != nil {
			ctx.AbortWithStatusJSON(400, fmt.Sprintf("Error: %v\n", err))
			log.Warning(fmt.Sprintf("Aborted with status 400: \nError: %v\n", err))
			return
		}
		ctx.JSON(http.StatusOK, service)
		log.Info("User is successfully checked.")
	}
}
