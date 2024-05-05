package main

import (
	"authentification/api"
	"authentification/dbconnect"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	dbconnect.ConnectDb()
	defer dbconnect.DataBase.Close(dbconnect.Ctx)
	dbconnect.CreateTable()
	adminRouter := gin.Default()
	adminRouter.POST("/insertUser", api.InsertHandler)
	adminRouter.POST("/deleteUser", api.DeleteUserHandler)
	adminRouter.POST("/addUserServices", api.UpdateServicesHandler)
	adminRouter.POST("/removeUserServices", api.RemoveServicesHanndler)
	port := os.Getenv("ADMIN_PORT")
	host := os.Getenv("ADMIN_HOST")
	url := fmt.Sprintf("%s:%s", host, port)
	adminRouter.Run(url)
}
