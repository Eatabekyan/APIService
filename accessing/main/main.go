package main

import (
	"accessing/api"
	"accessing/dbconnect"
	"fmt"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var (
	wg      sync.WaitGroup
	logFile = "/accessing/service_logs.log"
	logger  = log.New()
)

func adminServe() {
	adminRouter := gin.Default()
	adminRouter.POST("/insertUser", api.InsertHandler(logger))
	adminRouter.POST("/deleteUser", api.DeleteUserHandler(logger))
	adminRouter.POST("/addUserServices", api.AddServicesHandler(logger))
	adminRouter.POST("/removeUserServices", api.RemoveServicesHanndler(logger))
	port := os.Getenv("ADMIN_PORT")
	url := fmt.Sprintf(":%s", port)
	adminRouter.Run(url)
}

func clientServe() {
	clientRouter := gin.Default()
	clientRouter.POST("checkUserAccess", api.CheckUserAccessHandler(logger))
	port := os.Getenv("CLIENT_PORT")
	url := fmt.Sprintf(":%s", port)
	clientRouter.Run(url)
}

func ServePorts() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		adminServe()
	}()
	go func() {
		defer wg.Done()
		clientServe()
	}()
	wg.Wait()
}

func main() {
	logger.Out = os.Stdout
	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Warning(fmt.Sprintf("Could not open log file: %v\n", err))
	} else {
		defer f.Close()
		logger.Out = f
	}

	dbconnect.ConnectDb(logger)

	defer dbconnect.DataBase.Close(dbconnect.Ctx)
	dbconnect.CreateTable()
	ServePorts()
}
