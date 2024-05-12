package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	logger  = log.New()
	logFile = "/archive_manager/logs/archive_logs.log"
)

type ArchiveRequest struct {
	User          string `json:"user"`
	RequestRecord string `json:"record"`
}

type TokenResponse struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
	Scope     string `json:"scope"`
	Type string `json:"token_type"`
}

type CheckUserRequest struct {
	UserName string `json:"username"`
	Token    string `json:"token"`
}

func getToken() string {
	host := os.Getenv("OAUTH_SERVER_HOST")
	port, err := strconv.Atoi(os.Getenv("OAUTH_SERVER_PORT"))
	if err != nil {
		logger.Fatal(fmt.Sprintf("Error converting port: %v", err))
	}
	url := fmt.Sprintf("http://%s:%d/token", host, port)
	payload := "grant_type=client_credentials&client_id=111111&client_secret=888888"
	tokenreq, err := http.NewRequest("POST", url, bytes.NewBufferString(payload))
	if err != nil {
		logger.Warning(fmt.Sprintf("Error creating request: %v", err))
		return ""
	}
	tokenreq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(tokenreq)
	if err != nil {
		logger.Warning(fmt.Sprintf("Error making request:%v", err))
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Warning(fmt.Sprintf("Error reading response body:%v", err))
		return ""
	}
	responseStruct := &TokenResponse{}
	err = json.Unmarshal(body, &responseStruct)
	if err != nil {
		logger.Warning(fmt.Sprintf("Error unmarshalling response body:%v", err))
		return ""
	}
	return responseStruct.Token
}

func checkAccess(username, token string) (records []string) {
	host := os.Getenv("CLIENT_HOST")
	port, err := strconv.Atoi(os.Getenv("CLIENT_PORT"))
	if err != nil {
		logger.Fatal(fmt.Sprintf("Error converting port: %v", err))
	}
	url := fmt.Sprintf("http://%s:%d/checkUserAccess", host, port)
	accessReq := &CheckUserRequest{
		UserName: username,
		Token:    token,
	}
	payload, err := json.Marshal(accessReq)
	if err != nil {
		logger.Warning(fmt.Sprintf("Error creating request: %v", err))
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		logger.Warning(fmt.Sprintf("Error creating request: %v", err))
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Warning(fmt.Sprintf("Error making request:%v", err))
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Warning(fmt.Sprintf("Error reading response body:%v", err))
		return
	}
	logger.Info(fmt.Sprintf("Response Body: %s", string(body)))
	response := struct {
		Access []string `json:"access"`
	}{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		logger.Warning(fmt.Sprintf("Error unmarshalling response body:%v", err))
		return nil
	}
	records = response.Access
	return
}
func RequestArchiveHandler(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatusJSON(400, "User is not defined")
		logger.Warning(fmt.Sprintf("Aborted with status 400: User is not defined\nError: %v\n", err))
		return
	}
	req := ArchiveRequest{}
	err = json.Unmarshal(data, &req)
	if err != nil {
		c.AbortWithStatusJSON(400, "Bad Input")
		logger.Warning(fmt.Sprintf("Aborted with status 400: Bad Input:\nError: %v\n", err))
		return
	}
	token := getToken()
	if token == "" {
		return
	}
	records := checkAccess(req.User, token)

	for _, record := range records {
		if record == req.RequestRecord {
			c.JSON(http.StatusOK, "Accepted Access")
			logger.Info("Accepted Access")
			return
		}
	}
	c.AbortWithStatusJSON(400, "Access Denied")
	logger.Warning("Access Denied")
}
func main() {
	logger.Out = os.Stdout
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		logger.Warning(fmt.Sprintf("Could not open log file: %v\n", err))
	} else {
		defer f.Close()
		logger.Out = f
	}
	router := gin.Default()
	port := os.Getenv("ARCHIVE_PORT")
	router.POST("/requestArchive", RequestArchiveHandler)
	url := fmt.Sprintf(":%s", port)
	router.Run(url)

}
