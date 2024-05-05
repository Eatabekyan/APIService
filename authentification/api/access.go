package api

import (
	"authentification/dbtypes"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUserFromRequest(ctx *gin.Context) (*dbtypes.User, error) {
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(400, "User is not defined")
		log.Warning(fmt.Sprintf("Aborted with status 400: User is not defined\nError: %v\n", err))
		return nil, err
	}
	user := dbtypes.User{}
	err = json.Unmarshal(data, &user)
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Bad Input")
		log.Warning(fmt.Sprintf("Aborted with status 400: Bad Input:\nError: %v\n", err))
		return nil, err
	}
	return &user, nil
}
