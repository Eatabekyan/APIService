package dbconnect

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var DataBase *pgx.Conn = nil
var Ctx context.Context = context.Background()

func ConnectDb() error {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(fmt.Sprintf("Could not load environment: %v\n", err))
	}
	host := os.Getenv("DBHOST")
	port, _ := strconv.Atoi(os.Getenv("DBPORT"))
	user := os.Getenv("DBUSER")
	dbname := os.Getenv("DBNAME")
	pass := os.Getenv("DBPASSWORD")

	urlExample := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, pass, host, port, dbname)
	DataBase, err = pgx.Connect(Ctx, urlExample)

	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	if err != nil {
		log.Fatal(fmt.Sprintf("QueryRow failed: %v\n", err))
	}
	return err
}
