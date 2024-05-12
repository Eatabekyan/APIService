package api

import (
	"accessing/pkg/dbtypes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ResponseStruct struct {
	Access []string `json:"access"`
}

func getUserFromAdminRequest(log *log.Logger, ctx *gin.Context) (*dbtypes.User, error) {
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

func getClientRequestInfo(log *log.Logger, ctx *gin.Context) (*dbtypes.ClientRequest, error) {
	data, err := ctx.GetRawData()
	if err != nil {
		ctx.AbortWithStatusJSON(400, "User is not defined")
		log.Warning(fmt.Sprintf("Aborted with status 400: User is not defined\nError: %v\n", err))
		return nil, err
	}
	reqInfo := dbtypes.ClientRequest{}
	err = json.Unmarshal(data, &reqInfo)
	if err != nil {
		ctx.AbortWithStatusJSON(400, "Bad Input")
		log.Warning(fmt.Sprintf("Aborted with status 400: Bad Input:\nError: %v\n", err))
		return nil, err
	}
	return &reqInfo, nil
}

func checkToken(log *log.Logger, token string) (service string, err error) {
	oauthHost := os.Getenv("OAUTH_SERVER_HOST")
	oauthPort, err := strconv.Atoi(os.Getenv("OAUTH_SERVER_PORT"))
	if err != nil {
		log.Warning(fmt.Sprintf("Error : %v\n", err))
		return
	}
	myClientId := os.Getenv("MY_CLIENT_ID")
	mySecret := os.Getenv("MY_SECRET")
	url := fmt.Sprintf(
		"http://%s:%d/check-token?grant_type=client_credentials&client_id=%s&client_secret=%s",
		oauthHost,
		oauthPort,
		myClientId,
		mySecret,
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Warning(fmt.Sprintf("Error : %v\n", err))
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Warning(fmt.Sprintf("Error: %v\n", err))
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Warning(fmt.Sprintf("Error : %v\n", err))
		return
	}

	responseStruct := make(map[string]string, 0)

	err = json.Unmarshal(body, &responseStruct)
	if err != nil {
		return "", err
	}
	return responseStruct["service"], nil
}

func GetUserAccessToService(log *log.Logger, user *dbtypes.User, service string) ([]string, error) {
	data, err := json.Marshal(user.Access)
	if err != nil {
		return nil, err
	}
	access := map[string]map[string][]string{}
	err = json.Unmarshal(data, &access)
	if err != nil {
		return nil, err
	}
	if _, ok := access[service]; !ok {
		return nil, fmt.Errorf("no access to service %s", service)
	}
	return access[service]["access"], nil
}
